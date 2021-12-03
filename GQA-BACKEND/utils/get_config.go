package utils

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
)

func GetConfigBackend(gqaOption string) (key string) {
	var sysConfigBackend system.SysConfigBackend
	if err := global.GqaDb.Where("gqa_option = ?", gqaOption).First(&sysConfigBackend).Error; err != nil {
		return ""
	}
	if sysConfigBackend.Custom != "" {
		return sysConfigBackend.Custom
	} else {
		return sysConfigBackend.Default
	}
}

func GetConfigFrontend(gqaOption string) (key string) {
	var sysConfigFrontend system.SysConfigFrontend
	if err := global.GqaDb.Where("gqa_option = ?", gqaOption).First(&sysConfigFrontend).Error; err != nil {
		return ""
	}
	if sysConfigFrontend.Custom != "" {
		return sysConfigFrontend.Custom
	} else {
		return sysConfigFrontend.Default
	}
}
