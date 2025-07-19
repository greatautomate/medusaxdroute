package console_setting

import "medusaxd-api/setting/config"

type ConsoleSetting struct {
    ApiInfo           string `json:"api_info"`           // Console API information (JSON array string)
    UptimeKumaGroups  string `json:"uptime_kuma_groups"` // Uptime Kuma group configuration (JSON array string)
    Announcements     string `json:"announcements"`      // System announcements (JSON array string)
    FAQ               string `json:"faq"`                // Frequently asked questions (JSON array string)
    ApiInfoEnabled        bool `json:"api_info_enabled"`        // Whether to enable API information panel
    UptimeKumaEnabled     bool `json:"uptime_kuma_enabled"`     // Whether to enable Uptime Kuma panel
    AnnouncementsEnabled  bool `json:"announcements_enabled"`   // Whether to enable system announcements panel
    FAQEnabled            bool `json:"faq_enabled"`             // Whether to enable FAQ panel
}

// 默认配置
var defaultConsoleSetting = ConsoleSetting{
    ApiInfo:          "",
    UptimeKumaGroups: "",
    Announcements:    "",
    FAQ:              "",
    ApiInfoEnabled:       true,
    UptimeKumaEnabled:    true,
    AnnouncementsEnabled: true,
    FAQEnabled:           true,
}

// 全局实例
var consoleSetting = defaultConsoleSetting

func init() {
    // 注册到全局配置管理器，键名为 console_setting
    config.GlobalConfig.Register("console_setting", &consoleSetting)
}

// GetConsoleSetting 获取 ConsoleSetting 配置实例
func GetConsoleSetting() *ConsoleSetting {
    return &consoleSetting
} 
