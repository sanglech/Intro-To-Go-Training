package errhandle

import (
	"errors"
)

// Divide divides numerator by denominator (numerator / denominator) and returns the result.
//
// If denominator is 0, Divide returns an error (it doesn't matter what number you return in the error case).
func Divide(numerator, denominator float64) (float64, error) {
	if (denominator==0.0){
		return 0, errors.New("denominator can't be 0")
	}
	return numerator/denominator,nil
}

// Do not modify this type!
//
// UserDB is a persistent database containing information about users.
type UserDB interface {

	// GetUsername gets a user's username from storage filtering by their user ID.
	//
	// If a user with this ID is not found, returns NotFoundError.
	// If there is an unexpected error, returns that error.
	GetUsername(userID int) (username string, err error)
}

// Do not modify this type!
//
// UserCache is a fast but short lived KVS cache containing information about users.
type UserCache interface {

	// SetUsername sets a user's username in KVS cache for fast retrieval.
	SetUsername(userID int, username string) error
}

// Do not modify this variable!
var ErrNotFound = errors.New("user not found!")

// SetCacheFromDB attempts to select a User's username from DB and set the username in cache.
// This function is used for refreshing a cache when it has expired.
//
// SetCacheFromDB has three arguments:
//    1. userID - ID of the user to select from DB and insert into cache
//    2. db - A user database repository which can be used to get user info from the DB
//    3. cache - A user cache repository which can be used to set user info in the cache
//
// SetCacheFromDB has three results:
//    1. username - The username found in DB, if any. In an error case the value of this result is undefined.
//    2. userExists - True if the user was found in DB. False otherwise.
//    3. err - Err will not be nil only if an *unhandled* error occurs in SetCacheFromDB.
//
// This function should do the following:
//
//    Get info for userID from DB
//      * if user not found -> handle error (see comment about "err" result) and return that the user does not exist
//      * if unexpected error -> pass this unhandled error up the call stack for handling elsewhere
//
//    Set the info selected from DB into cache
//      * if unexpected error -> pass this unhandled error up the call stack for handling elsewhere
//
//    Return the username from DB & return that the user exists
func SetCacheFromDB(userID int, db UserDB, cache UserCache) (username string, userExists bool, err error) {
	dbUserName,errMsg:=db.GetUsername(userID)

	if errMsg==ErrNotFound {
		return "", false, nil
	}

	if errMsg != nil{
		return "", false, errMsg
	}

	cacheSetErr:= cache.SetUsername(userID,dbUserName)

	if cacheSetErr!=nil {
		return "", false, cacheSetErr
	}

	return dbUserName, true, nil //TODO implement
}
