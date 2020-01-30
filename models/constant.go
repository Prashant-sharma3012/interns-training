package models

import "time"

const (
	PASSWORD_MAXIMUM_LENGTH = 64
	PASSWORD_MINIMUM_LENGTH = 5

	TOKEN_MAX_AGE = time.Second * 60 * 24 * 7 // sec * min * hour * days

	API_URL_SUFFIX                = "/api/v1/"
	ALLOWED_FAILED_LOGIN_ATTEMPTS = 4
	ACCOUNT_LOCK_MINUTES          = 5

	ERRCODE1  = "USER_DOESNT_EXISTS"
	ERRCODE2  = "INVALID_LOGIN_ID"
	ERRCODE3  = "BAD_PASSWORD"
	ERRCODE4  = "EMAIL_ALREADY_REGISTERED"
	ERRCODE5  = "INVALID_DATA_OBJECT"
	ERRCODE6  = "BAD_DATA"
	ERRCODE7  = "INTERNAL_SERVER_ERROR"
	ERRCODE8  = "INVALID_ACCESS_TOKEN"
	ERRCODE9  = "ACCESS_TOKEN_MISSING"
	ERRCODE10 = "USER_ACCESS_BLOCKED"
	ERRCODE11 = "BAD_CREDENTIALS"
	ERRCODE12 = "USER_ACCESS_LOCKED"
	ERRCODE13 = "USER_NOT_AUTHORIZED"

	//user statuses
	USER_STATUS_ACTIVE  = "active"
	USER_STATUS_BLOCKED = "blocked"

	//admin statuses
	ADMIN_STATUS_ACTIVE  = "active"
	ADMIN_STATUS_BLOCKED = "blocked"

	//client statuses
	CLIENT_STATUS_ACTIVE  = "active"
	CLIENT_STATUS_BLOCKED = "blocked"

	//api response status
	FAILED_STATUS  = "failed"
	SUCCESS_STATUS = "success"

	USER_TYPE_ADMIN           = "admin"
	USER_TYPE_SITE_ENGINEER   = "siteEngineer"
	USER_TYPE_CLIENT          = "client"
	USER_TYPE_PROJECT_MANAGER = "projectManager"
	USER_TYPE_TENDERER        = "tenderer"
	USER_TYPE_ACCOUNTS        = "accounts"
	USER_TYPE_PROCUREMENT     = "procurement"
	USER_TYPE_ANY             = "any"

	INIT_ADMIN_USER_POINTS = 10

	//only admin privilage
	ACCESS_WRITE = "access_write"
	USER_WRITE   = "user_write"

	//other user privilages
	USER_READ         = "user_read"
	ITEM_READ         = "item_read"
	ITEM_WRITE        = "item_write"
	MASTER_ITEM_READ  = "master_item_read"
	MASTER_ITEM_WRITE = "master_item_write"
	PROJECT_WRITE     = "project_read"
	PROJECT_READ      = "project_write"
	BOQ_READ          = "boq_read"
	BOQ_WRITE         = "boq_write"
	CATEGORY_WRITE    = "category_write"
	CATEGORY_READ     = "category_read"
)
