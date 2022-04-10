package main

import (
	dogv1 "go.buf.build/grpc/go/petland/dogapis/petland/dog/v1"
	tracingv1 "go.buf.build/grpc/go/petland/tracingapis/petland/tracing/v1"
	date "google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/money"
	"google.golang.org/genproto/googleapis/type/postaladdress"
	"strings"
	"time"
)

func getBreed(breed string) dogv1.Breed {
	switch breed {
	case "GOLDEN_RETRIEVER":
		return dogv1.Breed_BREED_GOLDEN_RETRIEVER
	case "BREED_GREYHOUND":
		return dogv1.Breed_BREED_GREYHOUND
	case "BREED_PUG":
		return dogv1.Breed_BREED_PUG
	case "REED_BULLDOG":
		return dogv1.Breed_BREED_BULLDOG
	default:
		return dogv1.Breed_BREED_UNSPECIFIED
	}
}

func getBirthDate(unixTimeStamp int) *date.Date {
	t := time.Unix(int64(unixTimeStamp), 0)
	return &date.Date{
		Year:  int32(t.Year()),
		Month: int32(t.Month()),
		Day:   int32(t.Day()),
	}
}

func getMoney(cents int) *money.Money {
	return &money.Money{
		CurrencyCode: "USD",
		Units:        int64(cents / 100),
		Nanos:        int32(cents%100) * 10000000,
	}
}

func getPostalAddress(address string) *dogv1.Dog_PostalAddress {
	split := strings.Split(address, ",")
	for i, s := range split {
		split[i] = strings.TrimSpace(s)
	}

	return &dogv1.Dog_PostalAddress{
		PostalAddress: &postaladdress.PostalAddress{
			RegionCode:         "US",
			LanguageCode:       "en",
			PostalCode:         split[3],
			SortingCode:        split[2],
			AdministrativeArea: split[1],
			AddressLines:       []string{split[0]},
		},
	}
}

func getFavoriteFoods(favoriteFoods map[string]int) map[string]*money.Money {
	m := make(map[string]*money.Money)

	for key, element := range favoriteFoods {
		m[key] = getMoney(element)
	}

	return m
}

func GetResp(dog *Dog) dogv1.GetDogResponse {
	return dogv1.GetDogResponse{
		Dog: &dogv1.Dog{
			Id:                     dog.Id,
			Name:                   dog.Name,
			Age:                    dog.Age,
			Breed:                  getBreed(dog.Breed),
			BirthDate:              getBirthDate(dog.BirthDate),
			DogPrice:               getMoney(dog.DogPrice),
			FavoriteToys:           dog.FavoriteToys,
			ProfileImageUrl:        dog.ProfileImageUrl,
			IsVaccinated:           dog.IsVaccinated,
			Weight:                 dog.Weight,
			CollarInformation:      getPostalAddress(dog.PostalAddress),
			FavoriteFoodsWithPrice: getFavoriteFoods(dog.FavoriteFoodsWithPrice),
		},
		Trace: &tracingv1.Trace{
			TraceId: "v1",
		},
	}
}
