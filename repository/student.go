package repository

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/interns-training/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StudentRepo struct{}

func GetStudentRepo() *StudentRepo {
	return &StudentRepo{}
}

func (s *StudentRepo) List() ([]models.Student, error) {

	ctx := context.Background()
	studentCollection := DB.Collection("student")

	studentCur, err := studentCollection.Find(ctx, bson.M{}, options.Find())

	defer studentCur.Close(ctx)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var students []models.Student

	for studentCur.Next(nil) {

		student := models.Student{}
		err := studentCur.Decode(student)

		if err != nil {
			log.Fatal("Decode error ", err)
		}

		students = append(students, student)
	}

	return students, nil
}

func (s *StudentRepo) getStudentById(studentId string) (models.Student, error) {

	collection := DB.Collection("student")

	student := models.Student{}

	err := collection.FindOne(context.Background(), bson.D{{"id", studentId}}).Decode(&student)

	if err != nil {
		return student, err
	}

	return student, nil
}

func (s *StudentRepo) Add(student models.Student) ([]byte, error) {

	ctx := context.Background()
	collection := DB.Collection("student")

	res, err := collection.InsertOne(ctx, student)
	if err != nil {
		return nil, err
	}

	return []byte(res.InsertedID.(primitive.ObjectID).Hex()), nil
}

func (s *StudentRepo) Update(student models.Student) ([]byte, error) {

	ctx := context.Background()
	collection := DB.Collection("student")

	_, err := collection.UpdateOne(ctx, bson.D{{"id", student.Id}}, bson.D{{"$set", student}})

	if err != nil {
		return nil, err
	}

	return []byte(student.Id + "updated successfully"), nil
}

func (s *StudentRepo) Delete(student models.Student) ([]byte, error) {

	ctx := context.Background()
	collection := DB.Collection("student")

	res, err := collection.DeleteOne(ctx, bson.D{{"id", student.Id}})

	if err != nil {
		return nil, err
	}

	return []byte(strconv.Itoa(int(res.DeletedCount))), nil
}
