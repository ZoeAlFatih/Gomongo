package main

import (
	"fmt"
	"gomongo/config"
	"gomongo/modules/profile/model"
	"gomongo/modules/profile/repository"
	"time"
)

func main() {
	fmt.Println("Mongo Go")

	db, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")
	//saveProfile(profileRepository)
	//updateProfile(profileRepository)
	//deleteProfile(profileRepository)
	//getProfile("U2", profileRepository)
	getProfiles(profileRepository)
}

// Save Profile(profileRepository)
func saveProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "U3"
	p.FirsName = "Hasawao"
	p.LastName = "Bebas"
	p.Email = "hasawao@qwords.co.id"
	p.Password = "1234567"
	p.CreatedAt = time.Now()
	p.UpdateAt = time.Now()

	err := profileRepository.Save(&p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile saved..")
	}
}

// Update Profile(profileRepository)
func updateProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "U1"
	p.FirsName = "Hamdan"
	p.LastName = "Muhammad"
	p.Email = "hamdan@qwords.com"
	p.Password = "1234567890"
	p.CreatedAt = time.Now()
	p.UpdateAt = time.Now()

	err := profileRepository.Update("U1", &p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile update..")
	}
}

// Delete Profile(profileRepository)
func deleteProfile(profileRepository repository.ProfileRepository) {
	err := profileRepository.Delete("U1")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile deleted..")
	}

}

func getProfile(id string, profileRepository repository.ProfileRepository) {
	profile, err := profileRepository.FindByID(id)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(profile.ID)
	fmt.Println(profile.FirsName)
	fmt.Println(profile.LastName)
	fmt.Println(profile.Email)

}

func getProfiles(profileRepository repository.ProfileRepository) {
	profiles, err := profileRepository.FindAll()

	if err != nil {
		fmt.Println(err)
	}

	for _, profile := range profiles {
		fmt.Println("--------------------")
		fmt.Println(profile.ID)
		fmt.Println(profile.FirsName)
		fmt.Println(profile.LastName)
		fmt.Println(profile.Email)
	}
}
