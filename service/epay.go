package service

import (
	"medusaxd-api/setting"
)

func GetCallbackAddress() string {
	if setting.CustomCallbackAddress == "" {
		return setting.ServerAddress
	}
	return setting.CustomCallbackAddress
}
