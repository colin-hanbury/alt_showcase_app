package user

type UserInput struct {
	Name        string `bson:"name" json:"name"`
	Nationality string `bson:"nationality "json:"nationality"`
}
