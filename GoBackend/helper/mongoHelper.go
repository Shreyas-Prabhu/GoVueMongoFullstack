package helper

import (
	"context"
	"fmt"
	"log"

	"github.com/Shreyas-Prabhu/GoVueMongoFullstack.git/database"
	"github.com/Shreyas-Prabhu/GoVueMongoFullstack.git/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Collection = database.MongoClient.Database("MovieWorld").Collection("movies")

func AddMovie(movie model.Movie) string {
	_, err := Collection.InsertOne(context.TODO(), movie)
	if err != nil {
		panic(err)
	}
	return "Inserted"
}

func GetMovieById(id string) *model.Movie {
	movid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	var movie model.Movie
	filter := bson.M{"_id": movid}
	res := Collection.FindOne(context.Background(), filter).Decode(&movie)
	fmt.Println(res)
	return &movie
}

func GetAllMovies() *[]model.Movie {
	var movieArr []model.Movie
	filter := bson.D{{}}
	cursor, err := Collection.Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}

	for cursor.Next(context.Background()) {
		var movie model.Movie
		err := cursor.Decode(&movie)
		if err != nil {
			panic(err)
		}
		movieArr = append(movieArr, movie)
	}
	return &movieArr
}

func DeleteById(id string) string {
	movid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": movid}
	_, err = Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	return "Deleted"
}

func DeleteAll() string {
	filter := bson.D{{}}
	deleteResult, err := Collection.DeleteMany(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(deleteResult)
	return "Delete All"
}

func UpdateMovie(id string, movie model.Movie) *model.Movie {
	movid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": movid}
	update := bson.M{"$set": bson.M{"title": movie.Title, "year": movie.Year, "cast": movie.Cast, "genre": movie.Genre, "image": movie.Image}}
	updateResult, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println(updateResult)
	var updatedMovie model.Movie
	err = Collection.FindOne(context.Background(), filter).Decode(&updatedMovie)
	if err != nil {
		panic(err)
	}

	return &updatedMovie
}

func SearchResult(keyword string) *[]model.Movie {
	filter := bson.D{{"title", primitive.Regex{Pattern: keyword, Options: "i"}}}
	cursor, err := Collection.Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}

	var movieArr []model.Movie
	err = cursor.All(context.Background(), &movieArr)
	if err != nil {
		panic(err)
	}
	return &movieArr
}
