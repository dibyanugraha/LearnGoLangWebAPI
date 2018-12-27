package mongo_test

import (
	root "TestRestAPI/pkg"
	"TestRestAPI/pkg/mock"
	"TestRestAPI/pkg/mongo"
	"log"
	"testing"
)

const (
	mongoUrl           = "localhost:27017"
	dbName             = "test_db"
	userCollectionName = "user"
)

func Test_UserService(t *testing.T) {
	t.Run("CreateUser", createUserShouldInsertUserIntoMongo)
}

func createUserShouldInsertUserIntoMongo(t *testing.T) {
	// Arrange
	session, err := mongo.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	defer func() {
		session.DropDatabase(dbName)
		session.Close()
	}()

	mockHash := mock.Hash{}
	userService := mongo.NewUserService(session.Copy(), dbName, userCollectionName, &mockHash)

	testUserName := "integration_test_user"
	testPassword := "integration_test_password"

	user := root.User{
		Username: testUserName,
		Password: testPassword}

	//Act
	err = userService.Create(&user)

	// Assert
	if err != nil {
		t.Errorf("Unable to create user: `%s`", err)
	}
	var results []root.User
	session.GetCollection(dbName, userCollectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Error("Incorrect number of results. Expected `1`, got `%i`", count)
	}

	if results[0].Username != user.Username {
		t.Errorf("Incorrect username. Expected `%s`, got `%s`", testUserName, results[0].Username)
	}
}
