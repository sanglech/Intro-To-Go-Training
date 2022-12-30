package errhandle

import (
	"errors"
	"introduction-to-go-training/testutil"
	"testing"
)

// Tests for Divide

func TestDivide_DenominatorNotZero_ReturnsDivisionResult(t *testing.T) {
	var numer, denomin = 1.7, 8.8

	res, err := Divide(numer, denomin)

	testutil.AssertErrNil(t, err, true)
	testutil.AssertEqual(t, numer/denomin, res)
}

func TestDivide_DenominatorIsZero_ReturnsError(t *testing.T) {
	_, err := Divide(1.0, 0.0)

	testutil.AssertErrNil(t, err, false)
}

// Tests for SetCacheFromDB

type stubUserDB struct {
	GetUsernameFn func(userID int) (username string, err error)
}

func (s *stubUserDB) GetUsername(userID int) (username string, err error) {
	return s.GetUsernameFn(userID)
}

type stubUserCache struct {
	SetUsernameFn func(userID int, username string) error
}

func (s *stubUserCache) SetUsername(userID int, username string) error {
	return s.SetUsernameFn(userID, username)
}

func TestSetCacheFromDB_UsernameNotFound_ReturnsFalseForExists(t *testing.T) {
	testUserID := 123
	dbCalled := false
	db := &stubUserDB{GetUsernameFn: func(userID int) (string, error) {
		dbCalled = true
		testutil.AssertEqualMsg(t, testUserID, userID, "incorrect user ID passed to DB")
		return "", ErrNotFound
	}}
	cache := &stubUserCache{SetUsernameFn: func(userID int, username string) error {
		t.Fatal("cache should not be set when username is not found")
		return nil
	}}

	_, exists, err := SetCacheFromDB(testUserID, db, cache)

	testutil.AssertEqualMsg(t, true, dbCalled, "database should be selected from")
	testutil.AssertErrNil(t, err, true)
	testutil.AssertEqualMsg(t, false, exists, "userExists should be false when user is not found in DB")
}

func TestSetCacheFromDB_UnexpectedDBError_ReturnsError(t *testing.T) {
	testUserID := 123
	dbCalled := false
	expectedErr := errors.New("big problem")
	db := &stubUserDB{GetUsernameFn: func(userID int) (string, error) {
		dbCalled = true
		testutil.AssertEqualMsg(t, testUserID, userID, "incorrect user ID passed to DB")
		return "", expectedErr
	}}
	cache := &stubUserCache{SetUsernameFn: func(userID int, username string) error {
		t.Fatal("cache should not be set when DB error occurs")
		return nil
	}}

	_, _, err := SetCacheFromDB(testUserID, db, cache)

	testutil.AssertEqualMsg(t, true, dbCalled, "database should be selected from")
	testutil.AssertEqualMsg(t, true, errors.Is(err, expectedErr),
		"error from DB should be returned (optionally with stack trace)")
}

func TestSetCacheFromDB_UnexpectedCacheError_ReturnsError(t *testing.T) {
	testUserID := 123
	dbCalled := false
	cacheResult := "this is the username"
	db := &stubUserDB{GetUsernameFn: func(userID int) (string, error) {
		dbCalled = true
		testutil.AssertEqualMsg(t, testUserID, userID, "incorrect user ID passed to DB")
		return cacheResult, nil
	}}
	expectedErr := errors.New("big problem")
	cacheCalled := false
	cache := &stubUserCache{SetUsernameFn: func(userID int, username string) error {
		cacheCalled = true
		testutil.AssertEqualMsg(t, testUserID, userID, "incorrect user ID passed to cache")
		testutil.AssertEqualMsg(t, cacheResult, username, "DB result username should be input to cache")
		return expectedErr
	}}

	_, _, err := SetCacheFromDB(testUserID, db, cache)

	testutil.AssertEqualMsg(t, true, dbCalled, "database should be selected from")
	testutil.AssertEqualMsg(t, true, cacheCalled, "cache should be set")
	testutil.AssertEqualMsg(t, true, errors.Is(err, expectedErr),
		"error from cache should be returned (optionally with stack trace)")
}

func TestSetCacheFromDB_Success_ReturnsUsernameFromDBAfterSettingCache(t *testing.T) {
	testUserID := 123
	dbCalled := false
	dbResult := "this is the username"
	db := &stubUserDB{GetUsernameFn: func(userID int) (string, error) {
		dbCalled = true
		testutil.AssertEqualMsg(t, testUserID, userID, "incorrect user ID passed to DB")
		return dbResult, nil
	}}
	cacheCalled := false
	cache := &stubUserCache{SetUsernameFn: func(userID int, username string) error {
		cacheCalled = true
		testutil.AssertEqualMsg(t, testUserID, userID, "incorrect user ID passed to cache")
		testutil.AssertEqualMsg(t, dbResult, username, "DB result username should be input to cache")
		return nil
	}}

	res, exists, err := SetCacheFromDB(testUserID, db, cache)

	testutil.AssertEqualMsg(t, true, dbCalled, "database should be selected from")
	testutil.AssertEqualMsg(t, true, cacheCalled, "cache should be set")
	testutil.AssertEqualMsg(t, dbResult, res, "db selected username should be returned")
	testutil.AssertEqualMsg(t, true, exists, "user exists result should be true")
	testutil.AssertErrNil(t, err, true)
}
