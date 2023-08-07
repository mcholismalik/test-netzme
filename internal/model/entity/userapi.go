package entity

import "fmt"

type UserAPIEntity struct {
	Results []UserAPIEntityResult `json:"results"`
}

type UserAPIEntityResult struct {
	Gender   string                `json:"gender"`
	Name     UserAPIEntityName     `json:"name"`
	Location UserAPIEntityLocation `json:"location"`
	Picture  UserAPIEntityPicture  `json:"picture"`
}

type UserAPIEntityName struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type UserAPIEntityLocation struct {
	Street UserAPIEntityLocationStreet `json:"street"`
	City   string                      `json:"city"`
}

type UserAPIEntityLocationStreet struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
}

type UserAPIEntityPicture struct {
	Large string `json:"large"`
}

func (e UserAPIEntity) ToUserEntity() UserEntity {
	if len(e.Results) == 0 {
		return UserEntity{}
	}

	return UserEntity{
		Gender:   e.Results[0].Gender,
		Fullname: fmt.Sprintf(`%s %s %s`, e.Results[0].Name.Title, e.Results[0].Name.First, e.Results[0].Name.Last),
		Address:  fmt.Sprintf(`%d %s %s`, e.Results[0].Location.Street.Number, e.Results[0].Location.Street.Name, e.Results[0].Location.City),
		Picture:  e.Results[0].Picture.Large,
	}
}
