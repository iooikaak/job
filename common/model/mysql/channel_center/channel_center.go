package model

import (
	"time"
)

// AuthUser [...]
type AuthUser struct {
	ID      int64  `gorm:"primaryKey;column:id" json:"-"`
	Account string `gorm:"column:account" json:"account"`
	Name    string `gorm:"column:name" json:"name"`
	Pwd     string `gorm:"column:pwd" json:"pwd"`
}

// TableName get sql table name.获取数据库表名
func (m *AuthUser) TableName() string {
	return "auth_user"
}

// AuthUserColumns get sql column name.获取数据库列名
var AuthUserColumns = struct {
	ID      string
	Account string
	Name    string
	Pwd     string
}{
	ID:      "id",
	Account: "account",
	Name:    "name",
	Pwd:     "pwd",
}

// BgChannel [...]
type BgChannel struct {
	ChannelOID    string    `gorm:"primaryKey;column:channelOID" json:"-"`
	ChannelName   string    `gorm:"column:channelName" json:"channelName"`
	ChannelCode   string    `gorm:"column:channelCode" json:"channelCode"`
	ChannelFactor float64   `gorm:"column:channelFactor" json:"channelFactor"`
	VisitTimes    int       `gorm:"column:visitTimes" json:"visitTimes"`
	Comments      string    `gorm:"column:comments" json:"comments"`
	UpdateTime    time.Time `gorm:"column:updateTime" json:"updateTime"`
	CreateTime    time.Time `gorm:"column:createTime" json:"createTime"`
	IsDeleted     []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
	Enable        []uint8   `gorm:"column:enable" json:"enable"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannel) TableName() string {
	return "bg_channel"
}

// BgChannelColumns get sql column name.获取数据库列名
var BgChannelColumns = struct {
	ChannelOID    string
	ChannelName   string
	ChannelCode   string
	ChannelFactor string
	VisitTimes    string
	Comments      string
	UpdateTime    string
	CreateTime    string
	IsDeleted     string
	Enable        string
}{
	ChannelOID:    "channelOID",
	ChannelName:   "channelName",
	ChannelCode:   "channelCode",
	ChannelFactor: "channelFactor",
	VisitTimes:    "visitTimes",
	Comments:      "comments",
	UpdateTime:    "updateTime",
	CreateTime:    "createTime",
	IsDeleted:     "isDeleted",
	Enable:        "enable",
}

// BgChannelAuthorityControl [...]
type BgChannelAuthorityControl struct {
	ID             string    `gorm:"primaryKey;column:id" json:"-"`
	Channel        string    `gorm:"column:channel" json:"channel"`
	AuthorityKey   string    `gorm:"column:authority_key" json:"authorityKey"`
	AuthorityValue string    `gorm:"column:authority_value" json:"authorityValue"`
	IsDeleted      []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
	UpdateTime     time.Time `gorm:"column:updateTime" json:"updateTime"`
	CreateTime     time.Time `gorm:"column:createTime" json:"createTime"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelAuthorityControl) TableName() string {
	return "bg_channel_authority_control"
}

// BgChannelAuthorityControlColumns get sql column name.获取数据库列名
var BgChannelAuthorityControlColumns = struct {
	ID             string
	Channel        string
	AuthorityKey   string
	AuthorityValue string
	IsDeleted      string
	UpdateTime     string
	CreateTime     string
}{
	ID:             "id",
	Channel:        "channel",
	AuthorityKey:   "authority_key",
	AuthorityValue: "authority_value",
	IsDeleted:      "isDeleted",
	UpdateTime:     "updateTime",
	CreateTime:     "createTime",
}

// BgChannelCityMapping 渠道城市映射表
type BgChannelCityMapping struct {
	ID                       uint64                   `gorm:"primaryKey;column:id" json:"-"`
	LocalCityID              string                   `gorm:"column:localCityId" json:"localCityId"`                                                     // 系统城市Id
	LocalCityName            string                   `gorm:"column:localCityName" json:"localCityName"`                                                 // 系统城市名称
	ChannelCityID            string                   `gorm:"column:channelCityId" json:"channelCityId"`                                                 // 渠道城市id
	ChannelCityName          string                   `gorm:"column:channelCityName" json:"channelCityName"`                                             // 渠道城市名称
	ChannelName              string                   `gorm:"column:channelName" json:"channelName"`                                                     // 渠道名称
	ChannelProvinceMappingID uint64                   `gorm:"column:channelProvinceMappingId" json:"channelProvinceMappingId"`                           // 渠道省份映射id
	BgChannelProvinceMapping BgChannelProvinceMapping `gorm:"joinForeignKey:channelProvinceMappingId;foreignKey:id" json:"bgChannelProvinceMappingList"` // 渠道省份映射表
	Enable                   bool                     `gorm:"column:enable" json:"enable"`                                                               // 映射配置的状态 1有效，0无效
	CreateTime               time.Time                `gorm:"column:createTime" json:"createTime"`
	UpdateTime               time.Time                `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted                []uint8                  `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelCityMapping) TableName() string {
	return "bg_channel_city_mapping"
}

// BgChannelCityMappingColumns get sql column name.获取数据库列名
var BgChannelCityMappingColumns = struct {
	ID                       string
	LocalCityID              string
	LocalCityName            string
	ChannelCityID            string
	ChannelCityName          string
	ChannelName              string
	ChannelProvinceMappingID string
	Enable                   string
	CreateTime               string
	UpdateTime               string
	IsDeleted                string
}{
	ID:                       "id",
	LocalCityID:              "localCityId",
	LocalCityName:            "localCityName",
	ChannelCityID:            "channelCityId",
	ChannelCityName:          "channelCityName",
	ChannelName:              "channelName",
	ChannelProvinceMappingID: "channelProvinceMappingId",
	Enable:                   "enable",
	CreateTime:               "createTime",
	UpdateTime:               "updateTime",
	IsDeleted:                "isDeleted",
}

// BgChannelConfig 渠道配置表
type BgChannelConfig struct {
	ID          uint64    `gorm:"primaryKey;column:id" json:"-"`
	Code        string    `gorm:"column:code" json:"code"`               // 配置code
	Name        string    `gorm:"column:name" json:"name"`               // 配置名称
	Value       string    `gorm:"column:value" json:"value"`             // 配置值
	Ext1        string    `gorm:"column:ext1" json:"ext1"`               // 扩展参数1
	Ext2        string    `gorm:"column:ext2" json:"ext2"`               // 扩展参数2
	Channel     string    `gorm:"column:channel" json:"channel"`         // 渠道配置
	Description string    `gorm:"column:description" json:"description"` // 配置描述
	State       bool      `gorm:"column:state" json:"state"`             // 状态 1: 正常 0 异常
	CreateTime  time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime  time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted   []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelConfig) TableName() string {
	return "bg_channel_config"
}

// BgChannelConfigColumns get sql column name.获取数据库列名
var BgChannelConfigColumns = struct {
	ID          string
	Code        string
	Name        string
	Value       string
	Ext1        string
	Ext2        string
	Channel     string
	Description string
	State       string
	CreateTime  string
	UpdateTime  string
	IsDeleted   string
}{
	ID:          "id",
	Code:        "code",
	Name:        "name",
	Value:       "value",
	Ext1:        "ext1",
	Ext2:        "ext2",
	Channel:     "channel",
	Description: "description",
	State:       "state",
	CreateTime:  "createTime",
	UpdateTime:  "updateTime",
	IsDeleted:   "isDeleted",
}

// BgChannelContractCenter 供应商与分销商签约表
type BgChannelContractCenter struct {
	ID              int64     `gorm:"primaryKey;column:id" json:"-"`
	SupplierChannel string    `gorm:"column:supplier_channel" json:"supplierChannel"`  // 供应商渠道tenantId
	RetailChannel   string    `gorm:"column:retail_channel" json:"retailChannel"`      // 分销商渠道
	ChannelSystemID string    `gorm:"column:channel_system_id" json:"channelSystemId"` // 供应商渠道标识,如风火轮，则填写zbf，摩天轮则值为mtl
	NeedSynMessage  []uint8   `gorm:"column:need_syn_message" json:"needSynMessage"`   // 是否该渠道同步摩天轮消息
	CreateTime      time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime      time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted       []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelContractCenter) TableName() string {
	return "bg_channel_contract_center"
}

// BgChannelContractCenterColumns get sql column name.获取数据库列名
var BgChannelContractCenterColumns = struct {
	ID              string
	SupplierChannel string
	RetailChannel   string
	ChannelSystemID string
	NeedSynMessage  string
	CreateTime      string
	UpdateTime      string
	IsDeleted       string
}{
	ID:              "id",
	SupplierChannel: "supplier_channel",
	RetailChannel:   "retail_channel",
	ChannelSystemID: "channel_system_id",
	NeedSynMessage:  "need_syn_message",
	CreateTime:      "createTime",
	UpdateTime:      "updateTime",
	IsDeleted:       "isDeleted",
}

// BgChannelDataMappings [...]
type BgChannelDataMappings struct {
	ID           string    `gorm:"primaryKey;column:id" json:"-"`
	Type         string    `gorm:"column:type" json:"type"`                  // 数据类型
	Channel      string    `gorm:"column:channel" json:"channel"`            // 渠道名称
	MtlValue     string    `gorm:"column:mtl_value" json:"mtlValue"`         // 摩天轮数据
	ChannelValue string    `gorm:"column:channel_value" json:"channelValue"` // 渠道数据
	OutDescribe  string    `gorm:"column:out_describe" json:"outDescribe"`   // 字段补充
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted    []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelDataMappings) TableName() string {
	return "bg_channel_data_mappings"
}

// BgChannelDataMappingsColumns get sql column name.获取数据库列名
var BgChannelDataMappingsColumns = struct {
	ID           string
	Type         string
	Channel      string
	MtlValue     string
	ChannelValue string
	OutDescribe  string
	CreateTime   string
	UpdateTime   string
	IsDeleted    string
}{
	ID:           "id",
	Type:         "type",
	Channel:      "channel",
	MtlValue:     "mtl_value",
	ChannelValue: "channel_value",
	OutDescribe:  "out_describe",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
	IsDeleted:    "is_deleted",
}

// BgChannelDistrictMapping 渠道地区映射表
type BgChannelDistrictMapping struct {
	ID                       uint64                   `gorm:"primaryKey;column:id" json:"-"`
	LocalDistrictID          string                   `gorm:"column:localDistrictId" json:"localDistrictId"`                                             // 系统地区Id
	LocalDistrictName        string                   `gorm:"column:localDistrictName" json:"localDistrictName"`                                         // 系统地区名称
	ChannelDistrictID        string                   `gorm:"column:channelDistrictId" json:"channelDistrictId"`                                         // 渠道地区id
	ChannelDistrictName      string                   `gorm:"column:channelDistrictName" json:"channelDistrictName"`                                     // 渠道地区名称
	ChannelName              string                   `gorm:"column:channelName" json:"channelName"`                                                     // 渠道名称
	ChannelCityMappingID     uint64                   `gorm:"column:channelCityMappingId" json:"channelCityMappingId"`                                   // 渠道城市映射id
	BgChannelCityMapping     BgChannelCityMapping     `gorm:"joinForeignKey:channelCityMappingId;foreignKey:id" json:"bgChannelCityMappingList"`         // 渠道城市映射表
	ChannelProvinceMappingID uint64                   `gorm:"column:channelProvinceMappingId" json:"channelProvinceMappingId"`                           // 渠道省份映射id
	BgChannelProvinceMapping BgChannelProvinceMapping `gorm:"joinForeignKey:channelProvinceMappingId;foreignKey:id" json:"bgChannelProvinceMappingList"` // 渠道省份映射表
	Enable                   bool                     `gorm:"column:enable" json:"enable"`                                                               // 映射配置的状态 1有效，0无效
	CreateTime               time.Time                `gorm:"column:createTime" json:"createTime"`
	UpdateTime               time.Time                `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted                []uint8                  `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelDistrictMapping) TableName() string {
	return "bg_channel_district_mapping"
}

// BgChannelDistrictMappingColumns get sql column name.获取数据库列名
var BgChannelDistrictMappingColumns = struct {
	ID                       string
	LocalDistrictID          string
	LocalDistrictName        string
	ChannelDistrictID        string
	ChannelDistrictName      string
	ChannelName              string
	ChannelCityMappingID     string
	ChannelProvinceMappingID string
	Enable                   string
	CreateTime               string
	UpdateTime               string
	IsDeleted                string
}{
	ID:                       "id",
	LocalDistrictID:          "localDistrictId",
	LocalDistrictName:        "localDistrictName",
	ChannelDistrictID:        "channelDistrictId",
	ChannelDistrictName:      "channelDistrictName",
	ChannelName:              "channelName",
	ChannelCityMappingID:     "channelCityMappingId",
	ChannelProvinceMappingID: "channelProvinceMappingId",
	Enable:                   "enable",
	CreateTime:               "createTime",
	UpdateTime:               "updateTime",
	IsDeleted:                "isDeleted",
}

// BgChannelErrorReason [...]
type BgChannelErrorReason struct {
	ChannelErrorID string    `gorm:"primaryKey;column:channel_error_id" json:"-"`
	ResourceID     string    `gorm:"column:resource_id" json:"resourceId"`
	ResourceType   int       `gorm:"column:resource_type" json:"resourceType"` // 1:showOID;2:sessionOID;3:seatPlan
	Channel        string    `gorm:"column:channel" json:"channel"`
	ErrorReason    string    `gorm:"column:error_reason" json:"errorReason"`
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"updateTime"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelErrorReason) TableName() string {
	return "bg_channel_error_reason"
}

// BgChannelErrorReasonColumns get sql column name.获取数据库列名
var BgChannelErrorReasonColumns = struct {
	ChannelErrorID string
	ResourceID     string
	ResourceType   string
	Channel        string
	ErrorReason    string
	CreateTime     string
	UpdateTime     string
}{
	ChannelErrorID: "channel_error_id",
	ResourceID:     "resource_id",
	ResourceType:   "resource_type",
	Channel:        "channel",
	ErrorReason:    "error_reason",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
}

// BgChannelHistory tm_channel_history
type BgChannelHistory struct {
	ChannelHistoryOID string    `gorm:"primaryKey;column:channelHistoryOID" json:"-"`
	ChannelCode       string    `gorm:"column:channelCode" json:"channelCode"`
	ChannelLink       string    `gorm:"column:channelLink" json:"channelLink"`
	ChannelKey        string    `gorm:"column:channelKey" json:"channelKey"` // 渠道传过来的备注字段，方便以后添加渠道的区分，比如来源
	UpdateTime        time.Time `gorm:"column:updateTime" json:"updateTime"`
	CreateTime        time.Time `gorm:"column:createTime" json:"createTime"`
	IsDeleted         []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
	UserAgent         string    `gorm:"column:userAgent" json:"userAgent"`
	IP                string    `gorm:"column:ip" json:"ip"`
	Referrer          string    `gorm:"column:referrer" json:"referrer"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelHistory) TableName() string {
	return "bg_channel_history"
}

// BgChannelHistoryColumns get sql column name.获取数据库列名
var BgChannelHistoryColumns = struct {
	ChannelHistoryOID string
	ChannelCode       string
	ChannelLink       string
	ChannelKey        string
	UpdateTime        string
	CreateTime        string
	IsDeleted         string
	UserAgent         string
	IP                string
	Referrer          string
}{
	ChannelHistoryOID: "channelHistoryOID",
	ChannelCode:       "channelCode",
	ChannelLink:       "channelLink",
	ChannelKey:        "channelKey",
	UpdateTime:        "updateTime",
	CreateTime:        "createTime",
	IsDeleted:         "isDeleted",
	UserAgent:         "userAgent",
	IP:                "ip",
	Referrer:          "referrer",
}

// BgChannelInfo 渠道信息配置表
type BgChannelInfo struct {
	ID             int       `gorm:"primaryKey;column:id" json:"-"`
	ChannelCode    string    `gorm:"column:channelCode" json:"channelCode"`       // 渠道编号
	ChannelName    string    `gorm:"column:channelName" json:"channelName"`       // 渠道名称
	SellerID       string    `gorm:"column:sellerId" json:"sellerId"`             // 卖家id
	SupportExpress []uint8   `gorm:"column:supportExpress" json:"supportExpress"` // 渠道是否支持快递取票
	SupportOnsite  []uint8   `gorm:"column:supportOnsite" json:"supportOnsite"`   // 渠道是否支持上门取票
	SupportVenue   []uint8   `gorm:"column:supportVenue" json:"supportVenue"`     // 渠道是否支持现场取票
	CreateTime     time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime     time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted      []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelInfo) TableName() string {
	return "bg_channel_info"
}

// BgChannelInfoColumns get sql column name.获取数据库列名
var BgChannelInfoColumns = struct {
	ID             string
	ChannelCode    string
	ChannelName    string
	SellerID       string
	SupportExpress string
	SupportOnsite  string
	SupportVenue   string
	CreateTime     string
	UpdateTime     string
	IsDeleted      string
}{
	ID:             "id",
	ChannelCode:    "channelCode",
	ChannelName:    "channelName",
	SellerID:       "sellerId",
	SupportExpress: "supportExpress",
	SupportOnsite:  "supportOnsite",
	SupportVenue:   "supportVenue",
	CreateTime:     "createTime",
	UpdateTime:     "updateTime",
	IsDeleted:      "isDeleted",
}

// BgChannelJobSwitch 渠道定时任务开关
type BgChannelJobSwitch struct {
	JobName       string  `gorm:"primaryKey;column:job_name" json:"-"`        // 任务名称
	ExecuteSwitch []uint8 `gorm:"column:execute_switch" json:"executeSwitch"` // 执行开关， 0 关 1 开
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelJobSwitch) TableName() string {
	return "bg_channel_job_switch"
}

// BgChannelJobSwitchColumns get sql column name.获取数据库列名
var BgChannelJobSwitchColumns = struct {
	JobName       string
	ExecuteSwitch string
}{
	JobName:       "job_name",
	ExecuteSwitch: "execute_switch",
}

// BgChannelLink [...]
type BgChannelLink struct {
	ChannelLinkOID string    `gorm:"primaryKey;column:channelLinkOID" json:"-"`
	ChannelOID     string    `gorm:"column:channelOID" json:"channelOId"`
	ChannelLink    string    `gorm:"column:channelLink" json:"channelLink"`
	Comments       string    `gorm:"column:comments" json:"comments"`
	UpdateTime     time.Time `gorm:"column:updateTime" json:"updateTime"`
	CreateTime     time.Time `gorm:"column:createTime" json:"createTime"`
	IsDeleted      []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelLink) TableName() string {
	return "bg_channel_link"
}

// BgChannelLinkColumns get sql column name.获取数据库列名
var BgChannelLinkColumns = struct {
	ChannelLinkOID string
	ChannelOID     string
	ChannelLink    string
	Comments       string
	UpdateTime     string
	CreateTime     string
	IsDeleted      string
}{
	ChannelLinkOID: "channelLinkOID",
	ChannelOID:     "channelOID",
	ChannelLink:    "channelLink",
	Comments:       "comments",
	UpdateTime:     "updateTime",
	CreateTime:     "createTime",
	IsDeleted:      "isDeleted",
}

// BgChannelMessageQueue 渠道消息处理表
type BgChannelMessageQueue struct {
	ID           string    `gorm:"primaryKey;column:id" json:"-"`           // 主键
	ResourceType int8      `gorm:"column:resourceType" json:"resourceType"` // 消息类型(1.add  2.update  3.delete)
	Operation    int8      `gorm:"column:operation" json:"operation"`       // 操作类型(1 show,2 seatplan,3 venue,4 order_status,5 refund,9 order,6 show_status,7 seatplan_ticket_stock,8 session)
	ResourceID   string    `gorm:"column:resourceId" json:"resourceId"`     // 资源id
	SendTime     int64     `gorm:"column:sendTime" json:"sendTime"`         // 消息发出时间
	CreateTime   time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted    []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelMessageQueue) TableName() string {
	return "bg_channel_message_queue"
}

// BgChannelMessageQueueColumns get sql column name.获取数据库列名
var BgChannelMessageQueueColumns = struct {
	ID           string
	ResourceType string
	Operation    string
	ResourceID   string
	SendTime     string
	CreateTime   string
	UpdateTime   string
	IsDeleted    string
}{
	ID:           "id",
	ResourceType: "resourceType",
	Operation:    "operation",
	ResourceID:   "resourceId",
	SendTime:     "sendTime",
	CreateTime:   "createTime",
	UpdateTime:   "updateTime",
	IsDeleted:    "isDeleted",
}

// BgChannelMessageQueueBackup 渠道消息处理备份表
type BgChannelMessageQueueBackup struct {
	ID           string    `gorm:"column:id" json:"id"`                     // 主键
	ResourceType int8      `gorm:"column:resourceType" json:"resourceType"` // 消息类型(1.add  2.update  3.delete)
	Operation    int8      `gorm:"column:operation" json:"operation"`       // 操作类型(1 show,2 seatplan,3 venue,4 order_status,5 refund,9 order,6 show_status,7 seatplan_ticket_stock,8 session)
	ResourceID   string    `gorm:"column:resourceId" json:"resourceId"`     // 资源id
	SendTime     int64     `gorm:"column:sendTime" json:"sendTime"`         // 消息发出时间
	CreateTime   time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted    []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelMessageQueueBackup) TableName() string {
	return "bg_channel_message_queue_backup"
}

// BgChannelMessageQueueBackupColumns get sql column name.获取数据库列名
var BgChannelMessageQueueBackupColumns = struct {
	ID           string
	ResourceType string
	Operation    string
	ResourceID   string
	SendTime     string
	CreateTime   string
	UpdateTime   string
	IsDeleted    string
}{
	ID:           "id",
	ResourceType: "resourceType",
	Operation:    "operation",
	ResourceID:   "resourceId",
	SendTime:     "sendTime",
	CreateTime:   "createTime",
	UpdateTime:   "updateTime",
	IsDeleted:    "isDeleted",
}

// BgChannelMessageQueueCopy1 渠道消息处理表
type BgChannelMessageQueueCopy1 struct {
	ID           string    `gorm:"primaryKey;column:id" json:"-"`           // 主键
	ResourceType int8      `gorm:"column:resourceType" json:"resourceType"` // 消息类型(1.add  2.update  3.delete)
	Operation    int8      `gorm:"column:operation" json:"operation"`       // 操作类型(1 show,2 seatplan,3 venue,4 order_status,5 refund,9 order,6 show_status,7 seatplan_ticket_stock,8 session)
	ResourceID   string    `gorm:"column:resourceId" json:"resourceId"`     // 资源id
	SendTime     int64     `gorm:"column:sendTime" json:"sendTime"`         // 消息发出时间
	SentTime     int64     `gorm:"column:sentTime" json:"sentTime"`         // 消息发出时间
	CreateTime   time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted    []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelMessageQueueCopy1) TableName() string {
	return "bg_channel_message_queue_copy1"
}

// BgChannelMessageQueueCopy1Columns get sql column name.获取数据库列名
var BgChannelMessageQueueCopy1Columns = struct {
	ID           string
	ResourceType string
	Operation    string
	ResourceID   string
	SendTime     string
	SentTime     string
	CreateTime   string
	UpdateTime   string
	IsDeleted    string
}{
	ID:           "id",
	ResourceType: "resourceType",
	Operation:    "operation",
	ResourceID:   "resourceId",
	SendTime:     "sendTime",
	SentTime:     "sentTime",
	CreateTime:   "createTime",
	UpdateTime:   "updateTime",
	IsDeleted:    "isDeleted",
}

// BgChannelMqHandle 渠道供应商消息处理表
type BgChannelMqHandle struct {
	ID          int64     `gorm:"primaryKey;column:id" json:"-"`
	TimeStamp   int64     `gorm:"column:time_stamp" json:"timeStamp"`     // 消息时间戳
	ObjectID    string    `gorm:"column:object_id" json:"objectId"`       // 消息对象id
	Channel     string    `gorm:"column:channel" json:"channel"`          // 渠道
	OperateType string    `gorm:"column:operate_type" json:"operateType"` // 操作类型
	Type        string    `gorm:"column:type" json:"type"`                // 对象类型
	Finished    []uint8   `gorm:"column:finished" json:"finished"`        // 是否消息处理完成
	CreateTime  time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime  time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted   []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelMqHandle) TableName() string {
	return "bg_channel_mq_handle"
}

// BgChannelMqHandleColumns get sql column name.获取数据库列名
var BgChannelMqHandleColumns = struct {
	ID          string
	TimeStamp   string
	ObjectID    string
	Channel     string
	OperateType string
	Type        string
	Finished    string
	CreateTime  string
	UpdateTime  string
	IsDeleted   string
}{
	ID:          "id",
	TimeStamp:   "time_stamp",
	ObjectID:    "object_id",
	Channel:     "channel",
	OperateType: "operate_type",
	Type:        "type",
	Finished:    "finished",
	CreateTime:  "createTime",
	UpdateTime:  "updateTime",
	IsDeleted:   "isDeleted",
}

// BgChannelMtcTicketMappings [...]
type BgChannelMtcTicketMappings struct {
	ID                  string    `gorm:"primaryKey;column:id" json:"-"` // 映射Id
	ChannelTicketID     string    `gorm:"column:channel_ticket_id" json:"channelTicketId"`
	MtcTicketID         string    `gorm:"column:mtc_ticket_id" json:"mtcTicketId"` // MTC_SHOW_ID
	ChannelCode         string    `gorm:"column:channel_code" json:"channelCode"`  // 渠道编码
	LeftStock           int       `gorm:"column:left_stock" json:"leftStock"`      // 可售库存
	SupportSeat         []uint8   `gorm:"column:support_seat" json:"supportSeat"`  // 是否支持选座
	Row                 string    `gorm:"column:row" json:"row"`                   // 排号
	StartSeatNo         string    `gorm:"column:start_seat_no" json:"startSeatNo"` // 开始座位号
	EndSeatNo           string    `gorm:"column:end_seat_no" json:"endSeatNo"`     // 座位结束号
	CreateTime          time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime          time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted           []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
	MtcSeatPlanID       string    `gorm:"column:mtc_seat_plan_id" json:"mtcSeatPlanId"`
	MtcSectorConcreteID string    `gorm:"column:mtc_sector_concrete_id" json:"mtcSectorConcreteId"`
	MtcZoneConcreteID   string    `gorm:"column:mtc_zone_concrete_id" json:"mtcZoneConcreteId"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelMtcTicketMappings) TableName() string {
	return "bg_channel_mtc_ticket_mappings"
}

// BgChannelMtcTicketMappingsColumns get sql column name.获取数据库列名
var BgChannelMtcTicketMappingsColumns = struct {
	ID                  string
	ChannelTicketID     string
	MtcTicketID         string
	ChannelCode         string
	LeftStock           string
	SupportSeat         string
	Row                 string
	StartSeatNo         string
	EndSeatNo           string
	CreateTime          string
	UpdateTime          string
	IsDeleted           string
	MtcSeatPlanID       string
	MtcSectorConcreteID string
	MtcZoneConcreteID   string
}{
	ID:                  "id",
	ChannelTicketID:     "channel_ticket_id",
	MtcTicketID:         "mtc_ticket_id",
	ChannelCode:         "channel_code",
	LeftStock:           "left_stock",
	SupportSeat:         "support_seat",
	Row:                 "row",
	StartSeatNo:         "start_seat_no",
	EndSeatNo:           "end_seat_no",
	CreateTime:          "create_time",
	UpdateTime:          "update_time",
	IsDeleted:           "is_deleted",
	MtcSeatPlanID:       "mtc_seat_plan_id",
	MtcSectorConcreteID: "mtc_sector_concrete_id",
	MtcZoneConcreteID:   "mtc_zone_concrete_id",
}

// BgChannelOperationStore 操作标志记录
type BgChannelOperationStore struct {
	ResourceID    string    `gorm:"primaryKey;column:resource_id" json:"-"`     // 资源id
	OperationType int       `gorm:"primaryKey;column:operation_type" json:"-"`  // 操作类型号
	OperationFlag string    `gorm:"column:operation_flag" json:"operationFlag"` // 操作标记
	Channel       string    `gorm:"primaryKey;column:channel" json:"-"`         // 渠道名称
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted     []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelOperationStore) TableName() string {
	return "bg_channel_operation_store"
}

// BgChannelOperationStoreColumns get sql column name.获取数据库列名
var BgChannelOperationStoreColumns = struct {
	ResourceID    string
	OperationType string
	OperationFlag string
	Channel       string
	CreateTime    string
	UpdateTime    string
	IsDeleted     string
}{
	ResourceID:    "resource_id",
	OperationType: "operation_type",
	OperationFlag: "operation_flag",
	Channel:       "channel",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
	IsDeleted:     "is_deleted",
}

// BgChannelOrderDeliveryInfo [...]
type BgChannelOrderDeliveryInfo struct {
	ID                int       `gorm:"primaryKey;column:id" json:"-"`
	OrderNumber       string    `gorm:"column:orderNumber" json:"orderNumber"`
	ThirdOrederNumber string    `gorm:"column:thirdOrederNumber" json:"thirdOrederNumber"`
	Channel           string    `gorm:"column:channel" json:"channel"`
	UserName          string    `gorm:"column:userName" json:"userName"`
	UserPhone         string    `gorm:"column:userPhone" json:"userPhone"`
	SendMsg           []uint8   `gorm:"column:sendMsg" json:"sendMsg"` // 是否发送过短信
	IsDeleted         []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
	UpdateTime        time.Time `gorm:"column:update_time" json:"updateTime"`
	CreateTime        time.Time `gorm:"column:create_time" json:"createTime"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelOrderDeliveryInfo) TableName() string {
	return "bg_channel_order_delivery_info"
}

// BgChannelOrderDeliveryInfoColumns get sql column name.获取数据库列名
var BgChannelOrderDeliveryInfoColumns = struct {
	ID                string
	OrderNumber       string
	ThirdOrederNumber string
	Channel           string
	UserName          string
	UserPhone         string
	SendMsg           string
	IsDeleted         string
	UpdateTime        string
	CreateTime        string
}{
	ID:                "id",
	OrderNumber:       "orderNumber",
	ThirdOrederNumber: "thirdOrederNumber",
	Channel:           "channel",
	UserName:          "userName",
	UserPhone:         "userPhone",
	SendMsg:           "sendMsg",
	IsDeleted:         "is_deleted",
	UpdateTime:        "update_time",
	CreateTime:        "create_time",
}

// BgChannelOrderMapping 订单号对应表
type BgChannelOrderMapping struct {
	ID                  uint64    `gorm:"primaryKey;column:id" json:"-"`
	SupplierOrderNumber string    `gorm:"column:supplierOrderNumber" json:"supplierOrderNumber"` // 供应商订单号
	RetailerOrderNumber string    `gorm:"column:retailerOrderNumber" json:"retailerOrderNumber"` // 分销商订单号
	Retailer            string    `gorm:"column:retailer" json:"retailer"`                       // 分销商标识
	SupplierOrderStatus string    `gorm:"column:supplierOrderStatus" json:"supplierOrderStatus"` // 供应商订单状态
	RetailerOrderStatus string    `gorm:"column:retailerOrderStatus" json:"retailerOrderStatus"` // 分销商订单状态
	DeliveryType        string    `gorm:"column:deliveryType" json:"deliveryType"`               // 订单配送方式
	ExpressSynFlag      []uint8   `gorm:"column:expressSynFlag" json:"expressSynFlag"`           // 快递信息是否已同步
	Supplier            string    `gorm:"column:supplier" json:"supplier"`                       // 供应商标识
	RetryCount          int8      `gorm:"column:retryCount" json:"retryCount"`                   // 重试次数
	State               bool      `gorm:"column:state" json:"state"`                             // 订单状态 1: 正常 0 异常,2订单记录正常处理中
	ExtendInfo          string    `gorm:"column:extendInfo" json:"extendInfo"`                   // 扩展信息
	SupplierOrderID     string    `gorm:"column:supplierOrderId" json:"supplierOrderId"`         // 供应商订单Id
	RetailerOrderID     string    `gorm:"column:retailerOrderId" json:"retailerOrderId"`         // 分销商订单Id
	Reason              string    `gorm:"column:reason" json:"reason"`                           // 订单错误原因
	CreateTime          time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime          time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted           []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelOrderMapping) TableName() string {
	return "bg_channel_order_mapping"
}

// BgChannelOrderMappingColumns get sql column name.获取数据库列名
var BgChannelOrderMappingColumns = struct {
	ID                  string
	SupplierOrderNumber string
	RetailerOrderNumber string
	Retailer            string
	SupplierOrderStatus string
	RetailerOrderStatus string
	DeliveryType        string
	ExpressSynFlag      string
	Supplier            string
	RetryCount          string
	State               string
	ExtendInfo          string
	SupplierOrderID     string
	RetailerOrderID     string
	Reason              string
	CreateTime          string
	UpdateTime          string
	IsDeleted           string
}{
	ID:                  "id",
	SupplierOrderNumber: "supplierOrderNumber",
	RetailerOrderNumber: "retailerOrderNumber",
	Retailer:            "retailer",
	SupplierOrderStatus: "supplierOrderStatus",
	RetailerOrderStatus: "retailerOrderStatus",
	DeliveryType:        "deliveryType",
	ExpressSynFlag:      "expressSynFlag",
	Supplier:            "supplier",
	RetryCount:          "retryCount",
	State:               "state",
	ExtendInfo:          "extendInfo",
	SupplierOrderID:     "supplierOrderId",
	RetailerOrderID:     "retailerOrderId",
	Reason:              "reason",
	CreateTime:          "createTime",
	UpdateTime:          "updateTime",
	IsDeleted:           "isDeleted",
}

// BgChannelOrderPaymentRecord [...]
type BgChannelOrderPaymentRecord struct {
	ID            int       `gorm:"primaryKey;column:id" json:"-"`
	ThirdpartyOid string    `gorm:"column:thirdparty_oid" json:"thirdpartyOid"` // 第三方渠道订单号
	OrderID       string    `gorm:"column:order_id" json:"orderId"`             // 摩天轮订单号
	Channel       string    `gorm:"column:channel" json:"channel"`
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted     []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelOrderPaymentRecord) TableName() string {
	return "bg_channel_order_payment_record"
}

// BgChannelOrderPaymentRecordColumns get sql column name.获取数据库列名
var BgChannelOrderPaymentRecordColumns = struct {
	ID            string
	ThirdpartyOid string
	OrderID       string
	Channel       string
	CreateTime    string
	UpdateTime    string
	IsDeleted     string
}{
	ID:            "id",
	ThirdpartyOid: "thirdparty_oid",
	OrderID:       "order_id",
	Channel:       "channel",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
	IsDeleted:     "is_deleted",
}

// BgChannelOrderRecord 渠道订单补单记录表
type BgChannelOrderRecord struct {
	Channel        string    `gorm:"primaryKey;column:channel" json:"-"`          // 渠道
	ChannelOrderID string    `gorm:"primaryKey;column:channel_order_id" json:"-"` // 渠道订单id
	OrderID        string    `gorm:"column:order_id" json:"orderId"`              // 平台订单id
	AddOrderState  bool      `gorm:"column:add_order_state" json:"addOrderState"` // 是否补单成功 0 待补单 1 补单成功 2 补单不成功
	FailReason     string    `gorm:"column:fail_reason" json:"failReason"`        // 补单不成功原因
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted      []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelOrderRecord) TableName() string {
	return "bg_channel_order_record"
}

// BgChannelOrderRecordColumns get sql column name.获取数据库列名
var BgChannelOrderRecordColumns = struct {
	Channel        string
	ChannelOrderID string
	OrderID        string
	AddOrderState  string
	FailReason     string
	CreateTime     string
	UpdateTime     string
	IsDeleted      string
}{
	Channel:        "channel",
	ChannelOrderID: "channel_order_id",
	OrderID:        "order_id",
	AddOrderState:  "add_order_state",
	FailReason:     "fail_reason",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
	IsDeleted:      "is_deleted",
}

// BgChannelOrderSmsInfo [...]
type BgChannelOrderSmsInfo struct {
	ID              int       `gorm:"primaryKey;column:id" json:"-"`
	OrderID         string    `gorm:"column:order_id" json:"orderId"`
	OrderNumber     string    `gorm:"column:order_number" json:"orderNumber"`
	LinkURL         string    `gorm:"column:link_url" json:"linkUrl"`                // 短信链接
	SellerNickName  string    `gorm:"column:seller_nick_name" json:"sellerNickName"` // 卖家昵称
	ArriveTime      string    `gorm:"column:arrive_time" json:"arriveTime"`
	SellerCellPhone string    `gorm:"column:seller_cell_phone" json:"sellerCellPhone"`
	Address         string    `gorm:"column:address" json:"address"`
	Lng             string    `gorm:"column:lng" json:"lng"`
	Lat             string    `gorm:"column:lat" json:"lat"`
	CheckCode       string    `gorm:"column:check_code" json:"checkCode"` // 取票码
	Channel         string    `gorm:"column:channel" json:"channel"`
	Msg             string    `gorm:"column:msg" json:"msg"`
	Version         string    `gorm:"column:version" json:"version"`              // 版本
	SupplementMsg   string    `gorm:"column:supplement_msg" json:"supplementMsg"` // 补充短信内容
	UpdateTime      time.Time `gorm:"column:update_time" json:"updateTime"`
	CreateTime      time.Time `gorm:"column:create_time" json:"createTime"`
	IsDeleted       []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
	NotifyChannelID string    `gorm:"column:notify_channel_id" json:"notifyChannelId"` // 通知中心的productChannelId
	TemplateCode    string    `gorm:"column:template_code" json:"templateCode"`        // 通知短信模板编码
	SmsType         string    `gorm:"column:sms_type" json:"smsType"`                  // 短信类型
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelOrderSmsInfo) TableName() string {
	return "bg_channel_order_sms_info"
}

// BgChannelOrderSmsInfoColumns get sql column name.获取数据库列名
var BgChannelOrderSmsInfoColumns = struct {
	ID              string
	OrderID         string
	OrderNumber     string
	LinkURL         string
	SellerNickName  string
	ArriveTime      string
	SellerCellPhone string
	Address         string
	Lng             string
	Lat             string
	CheckCode       string
	Channel         string
	Msg             string
	Version         string
	SupplementMsg   string
	UpdateTime      string
	CreateTime      string
	IsDeleted       string
	NotifyChannelID string
	TemplateCode    string
	SmsType         string
}{
	ID:              "id",
	OrderID:         "order_id",
	OrderNumber:     "order_number",
	LinkURL:         "link_url",
	SellerNickName:  "seller_nick_name",
	ArriveTime:      "arrive_time",
	SellerCellPhone: "seller_cell_phone",
	Address:         "address",
	Lng:             "lng",
	Lat:             "lat",
	CheckCode:       "check_code",
	Channel:         "channel",
	Msg:             "msg",
	Version:         "version",
	SupplementMsg:   "supplement_msg",
	UpdateTime:      "update_time",
	CreateTime:      "create_time",
	IsDeleted:       "is_deleted",
	NotifyChannelID: "notify_channel_id",
	TemplateCode:    "template_code",
	SmsType:         "sms_type",
}

// BgChannelPostMappings [...]
type BgChannelPostMappings struct {
	MappingID     string    `gorm:"primaryKey;column:mapping_id" json:"-"`
	Channel       string    `gorm:"column:channel" json:"channel"`
	ChannelPostID string    `gorm:"column:channel_post_id" json:"channelPostId"`
	LikeCount     int       `gorm:"column:like_count" json:"likeCount"`
	Version       string    `gorm:"column:version" json:"version"`
	IsDeleted     []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelPostMappings) TableName() string {
	return "bg_channel_post_mappings"
}

// BgChannelPostMappingsColumns get sql column name.获取数据库列名
var BgChannelPostMappingsColumns = struct {
	MappingID     string
	Channel       string
	ChannelPostID string
	LikeCount     string
	Version       string
	IsDeleted     string
	UpdateTime    string
	CreateTime    string
}{
	MappingID:     "mapping_id",
	Channel:       "channel",
	ChannelPostID: "channel_post_id",
	LikeCount:     "like_count",
	Version:       "version",
	IsDeleted:     "is_deleted",
	UpdateTime:    "update_time",
	CreateTime:    "create_time",
}

// BgChannelPremiumCoefficient [...]
type BgChannelPremiumCoefficient struct {
	ID                      int     `gorm:"primaryKey;column:id" json:"-"`
	Channel                 string  `gorm:"column:channel" json:"channel"`
	PremiumRatio            float64 `gorm:"column:premiumRatio" json:"premiumRatio"`                       // 溢价系数
	SeatPickingPremiumRatio float64 `gorm:"column:seatPickingPremiumRatio" json:"seatPickingPremiumRatio"` // 选座票溢价系数
	IsDeleted               []uint8 `gorm:"column:isDeleted" json:"isDeleted"`
	OptimizeTicketType      string  `gorm:"column:optimize_ticket_type" json:"optimizeTicketType"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelPremiumCoefficient) TableName() string {
	return "bg_channel_premium_coefficient"
}

// BgChannelPremiumCoefficientColumns get sql column name.获取数据库列名
var BgChannelPremiumCoefficientColumns = struct {
	ID                      string
	Channel                 string
	PremiumRatio            string
	SeatPickingPremiumRatio string
	IsDeleted               string
	OptimizeTicketType      string
}{
	ID:                      "id",
	Channel:                 "channel",
	PremiumRatio:            "premiumRatio",
	SeatPickingPremiumRatio: "seatPickingPremiumRatio",
	IsDeleted:               "isDeleted",
	OptimizeTicketType:      "optimize_ticket_type",
}

// BgChannelPriceCoefficient 渠道价格系数表
type BgChannelPriceCoefficient struct {
	ID                 int       `gorm:"primaryKey;column:id" json:"-"`          // 主键
	ChannelCode        string    `gorm:"column:channel_code" json:"channelCode"` // 渠道编号
	PriceCoefficient   float64   `gorm:"column:price_coefficient" json:"priceCoefficient"`
	OptimizeTicketType string    `gorm:"column:optimize_ticket_type" json:"optimizeTicketType"`
	CreateTime         time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime         time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted          []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelPriceCoefficient) TableName() string {
	return "bg_channel_price_coefficient"
}

// BgChannelPriceCoefficientColumns get sql column name.获取数据库列名
var BgChannelPriceCoefficientColumns = struct {
	ID                 string
	ChannelCode        string
	PriceCoefficient   string
	OptimizeTicketType string
	CreateTime         string
	UpdateTime         string
	IsDeleted          string
}{
	ID:                 "id",
	ChannelCode:        "channel_code",
	PriceCoefficient:   "price_coefficient",
	OptimizeTicketType: "optimize_ticket_type",
	CreateTime:         "create_time",
	UpdateTime:         "update_time",
	IsDeleted:          "is_deleted",
}

// BgChannelPriceStock 渠道推送价格和库存表
type BgChannelPriceStock struct {
	ID         string    `gorm:"primaryKey;column:id" json:"-"`       // 主键
	ResourceID string    `gorm:"column:resourceId" json:"resourceId"` // 资源id，一般为票面id
	Price      float64   `gorm:"column:price" json:"price"`           // 渠道推送价格
	Stock      int       `gorm:"column:stock" json:"stock"`           // 推送的库存
	Channel    string    `gorm:"column:channel" json:"channel"`       // 渠道
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted  []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelPriceStock) TableName() string {
	return "bg_channel_price_stock"
}

// BgChannelPriceStockColumns get sql column name.获取数据库列名
var BgChannelPriceStockColumns = struct {
	ID         string
	ResourceID string
	Price      string
	Stock      string
	Channel    string
	CreateTime string
	UpdateTime string
	IsDeleted  string
}{
	ID:         "id",
	ResourceID: "resourceId",
	Price:      "price",
	Stock:      "stock",
	Channel:    "channel",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	IsDeleted:  "is_deleted",
}

// BgChannelPriceStockCopy1 渠道推送价格和库存表
type BgChannelPriceStockCopy1 struct {
	ID         string    `gorm:"primaryKey;column:id" json:"-"`       // 主键
	ResourceID string    `gorm:"column:resourceId" json:"resourceId"` // 资源id，一般为票面id
	Price      float64   `gorm:"column:price" json:"price"`           // 渠道推送价格
	Stock      int       `gorm:"column:stock" json:"stock"`           // 推送的库存
	Channel    string    `gorm:"column:channel" json:"channel"`       // 渠道
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted  []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelPriceStockCopy1) TableName() string {
	return "bg_channel_price_stock_copy1"
}

// BgChannelPriceStockCopy1Columns get sql column name.获取数据库列名
var BgChannelPriceStockCopy1Columns = struct {
	ID         string
	ResourceID string
	Price      string
	Stock      string
	Channel    string
	CreateTime string
	UpdateTime string
	IsDeleted  string
}{
	ID:         "id",
	ResourceID: "resourceId",
	Price:      "price",
	Stock:      "stock",
	Channel:    "channel",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	IsDeleted:  "is_deleted",
}

// BgChannelProvinceMapping 渠道省份映射表
type BgChannelProvinceMapping struct {
	ID                  uint64    `gorm:"primaryKey;column:id" json:"-"`
	LocalProvinceID     string    `gorm:"column:localProvinceId" json:"localProvinceId"`         // 系统省份Id
	LocalProvinceName   string    `gorm:"column:localProvinceName" json:"localProvinceName"`     // 系统省份名称
	ChannelProvinceID   string    `gorm:"column:channelProvinceId" json:"channelProvinceId"`     // 渠道省份id
	ChannelProvinceName string    `gorm:"column:channelProvinceName" json:"channelProvinceName"` // 渠道省份名称
	ChannelName         string    `gorm:"column:channelName" json:"channelName"`                 // 渠道名称
	Enable              bool      `gorm:"column:enable" json:"enable"`                           // 映射配置的状态 1有效，0无效
	CreateTime          time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime          time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted           []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelProvinceMapping) TableName() string {
	return "bg_channel_province_mapping"
}

// BgChannelProvinceMappingColumns get sql column name.获取数据库列名
var BgChannelProvinceMappingColumns = struct {
	ID                  string
	LocalProvinceID     string
	LocalProvinceName   string
	ChannelProvinceID   string
	ChannelProvinceName string
	ChannelName         string
	Enable              string
	CreateTime          string
	UpdateTime          string
	IsDeleted           string
}{
	ID:                  "id",
	LocalProvinceID:     "localProvinceId",
	LocalProvinceName:   "localProvinceName",
	ChannelProvinceID:   "channelProvinceId",
	ChannelProvinceName: "channelProvinceName",
	ChannelName:         "channelName",
	Enable:              "enable",
	CreateTime:          "createTime",
	UpdateTime:          "updateTime",
	IsDeleted:           "isDeleted",
}

// BgChannelPushedShow 各渠道已推送演出记录
type BgChannelPushedShow struct {
	ResourceID string    `gorm:"primaryKey;column:resource_id" json:"-"` // 资源id
	Type       int8      `gorm:"primaryKey;column:type" json:"-"`        // 资源类型
	Channel    string    `gorm:"primaryKey;column:channel" json:"-"`     // 渠道编号
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted  bool      `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelPushedShow) TableName() string {
	return "bg_channel_pushed_show"
}

// BgChannelPushedShowColumns get sql column name.获取数据库列名
var BgChannelPushedShowColumns = struct {
	ResourceID string
	Type       string
	Channel    string
	CreateTime string
	UpdateTime string
	IsDeleted  string
}{
	ResourceID: "resource_id",
	Type:       "type",
	Channel:    "channel",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	IsDeleted:  "is_deleted",
}

// BgChannelQueryValidateType [...]
type BgChannelQueryValidateType struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`
	Channel    string    `gorm:"column:channel" json:"channel"`
	TypeValue  string    `gorm:"column:type_value" json:"typeValue"`
	TypeKey    string    `gorm:"column:type_key" json:"typeKey"`
	IsDeleted  []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
	UpdateTime time.Time `gorm:"column:updateTime" json:"updateTime"`
	CreateTime time.Time `gorm:"column:createTime" json:"createTime"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelQueryValidateType) TableName() string {
	return "bg_channel_query_validate_type"
}

// BgChannelQueryValidateTypeColumns get sql column name.获取数据库列名
var BgChannelQueryValidateTypeColumns = struct {
	ID         string
	Channel    string
	TypeValue  string
	TypeKey    string
	IsDeleted  string
	UpdateTime string
	CreateTime string
}{
	ID:         "id",
	Channel:    "channel",
	TypeValue:  "type_value",
	TypeKey:    "type_key",
	IsDeleted:  "isDeleted",
	UpdateTime: "updateTime",
	CreateTime: "createTime",
}

// BgChannelRegionRelateHistory 供应商区域关联历史操作表
type BgChannelRegionRelateHistory struct {
	ID                   string    `gorm:"primaryKey;column:id" json:"-"`
	ChannelShowMappingID string    `gorm:"column:channel_show_mapping_id" json:"channelShowMappingId"` // 演出映射ID
	Content              string    `gorm:"column:content" json:"content"`                              // 变更内容
	RecordStatus         string    `gorm:"column:record_status" json:"recordStatus"`                   // 变更后状态
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`
	Operator             string    `gorm:"column:operator" json:"operator"`          // 操作人ID
	OperatorName         string    `gorm:"column:operator_name" json:"operatorName"` // 操作人名字
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelRegionRelateHistory) TableName() string {
	return "bg_channel_region_relate_history"
}

// BgChannelRegionRelateHistoryColumns get sql column name.获取数据库列名
var BgChannelRegionRelateHistoryColumns = struct {
	ID                   string
	ChannelShowMappingID string
	Content              string
	RecordStatus         string
	CreateTime           string
	Operator             string
	OperatorName         string
}{
	ID:                   "id",
	ChannelShowMappingID: "channel_show_mapping_id",
	Content:              "content",
	RecordStatus:         "record_status",
	CreateTime:           "create_time",
	Operator:             "operator",
	OperatorName:         "operator_name",
}

// BgChannelSeatPosition [...]
type BgChannelSeatPosition struct {
	ID           int64     `gorm:"primaryKey;column:id" json:"-"`
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`
	IsDeleted    []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`
	ColumnNumber int       `gorm:"column:column_number" json:"columnNumber"`
	OrderNumber  string    `gorm:"column:order_number" json:"orderNumber"`
	RegionName   string    `gorm:"column:region_name" json:"regionName"`
	RowNumber    int       `gorm:"column:row_number" json:"rowNumber"`
	SeatPlanID   string    `gorm:"column:seat_plan_id" json:"seatPlanId"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelSeatPosition) TableName() string {
	return "bg_channel_seat_position"
}

// BgChannelSeatPositionColumns get sql column name.获取数据库列名
var BgChannelSeatPositionColumns = struct {
	ID           string
	CreateTime   string
	IsDeleted    string
	UpdateTime   string
	ColumnNumber string
	OrderNumber  string
	RegionName   string
	RowNumber    string
	SeatPlanID   string
}{
	ID:           "id",
	CreateTime:   "create_time",
	IsDeleted:    "is_deleted",
	UpdateTime:   "update_time",
	ColumnNumber: "column_number",
	OrderNumber:  "order_number",
	RegionName:   "region_name",
	RowNumber:    "row_number",
	SeatPlanID:   "seat_plan_id",
}

// BgChannelSellerBackUp [...]
type BgChannelSellerBackUp struct {
	ID           int64     `gorm:"primaryKey;column:id" json:"-"`
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`
	IsDeleted    []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`
	RegisterTime time.Time `gorm:"column:register_time" json:"registerTime"`
	SellerID     string    `gorm:"column:seller_id" json:"sellerId"`
	System       string    `gorm:"column:system" json:"system"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelSellerBackUp) TableName() string {
	return "bg_channel_seller_back_up"
}

// BgChannelSellerBackUpColumns get sql column name.获取数据库列名
var BgChannelSellerBackUpColumns = struct {
	ID           string
	CreateTime   string
	IsDeleted    string
	UpdateTime   string
	RegisterTime string
	SellerID     string
	System       string
}{
	ID:           "id",
	CreateTime:   "create_time",
	IsDeleted:    "is_deleted",
	UpdateTime:   "update_time",
	RegisterTime: "register_time",
	SellerID:     "seller_id",
	System:       "system",
}

// BgChannelSellerTenantMapping 卖家租户id映射表
type BgChannelSellerTenantMapping struct {
	ID           uint64    `gorm:"primaryKey;column:id" json:"-"`
	SellerID     string    `gorm:"column:sellerId" json:"sellerId"`         // 卖家id
	TenantID     string    `gorm:"column:tenantId" json:"tenantId"`         // 租户id
	SupplierType string    `gorm:"column:supplierType" json:"supplierType"` // 供应商类别
	Enable       bool      `gorm:"column:enable" json:"enable"`             // 映射配置的状态 1有效，0无效
	SynFlag      []uint8   `gorm:"column:synFlag" json:"synFlag"`           // 状态是否已同步 ,当映射状态为false时，需将该状态同步到卖家端。1：已同步 0：未同步
	Version      int64     `gorm:"column:version" json:"version"`           // 版本号,与supplier_seat_plan相同
	Remark       string    `gorm:"column:remark" json:"remark"`             // 映射配置的备注
	CreateTime   time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted    []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelSellerTenantMapping) TableName() string {
	return "bg_channel_seller_tenant_mapping"
}

// BgChannelSellerTenantMappingColumns get sql column name.获取数据库列名
var BgChannelSellerTenantMappingColumns = struct {
	ID           string
	SellerID     string
	TenantID     string
	SupplierType string
	Enable       string
	SynFlag      string
	Version      string
	Remark       string
	CreateTime   string
	UpdateTime   string
	IsDeleted    string
}{
	ID:           "id",
	SellerID:     "sellerId",
	TenantID:     "tenantId",
	SupplierType: "supplierType",
	Enable:       "enable",
	SynFlag:      "synFlag",
	Version:      "version",
	Remark:       "remark",
	CreateTime:   "createTime",
	UpdateTime:   "updateTime",
	IsDeleted:    "isDeleted",
}

// BgChannelSellerTenantMappingCopy1 卖家租户id映射表
type BgChannelSellerTenantMappingCopy1 struct {
	ID           uint64    `gorm:"primaryKey;column:id" json:"-"`
	SellerID     string    `gorm:"column:sellerId" json:"sellerId"`         // 卖家id
	TenantID     string    `gorm:"column:tenantId" json:"tenantId"`         // 租户id
	SupplierType string    `gorm:"column:supplierType" json:"supplierType"` // 供应商类别
	Enable       bool      `gorm:"column:enable" json:"enable"`             // 映射配置的状态 1有效，0无效
	SynFlag      []uint8   `gorm:"column:synFlag" json:"synFlag"`           // 状态是否已同步 ,当映射状态为false时，需将该状态同步到卖家端。1：已同步 0：未同步
	Remark       string    `gorm:"column:remark" json:"remark"`             // 映射配置的备注
	CreateTime   time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted    []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelSellerTenantMappingCopy1) TableName() string {
	return "bg_channel_seller_tenant_mapping_copy1"
}

// BgChannelSellerTenantMappingCopy1Columns get sql column name.获取数据库列名
var BgChannelSellerTenantMappingCopy1Columns = struct {
	ID           string
	SellerID     string
	TenantID     string
	SupplierType string
	Enable       string
	SynFlag      string
	Remark       string
	CreateTime   string
	UpdateTime   string
	IsDeleted    string
}{
	ID:           "id",
	SellerID:     "sellerId",
	TenantID:     "tenantId",
	SupplierType: "supplierType",
	Enable:       "enable",
	SynFlag:      "synFlag",
	Remark:       "remark",
	CreateTime:   "createTime",
	UpdateTime:   "updateTime",
	IsDeleted:    "isDeleted",
}

// BgChannelShow [...]
type BgChannelShow struct {
	ShowID     string    `gorm:"column:show_id" json:"showId"`  // 资源id
	Channel    string    `gorm:"column:channel" json:"channel"` // 渠道编号
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted  []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelShow) TableName() string {
	return "bg_channel_show"
}

// BgChannelShowColumns get sql column name.获取数据库列名
var BgChannelShowColumns = struct {
	ShowID     string
	Channel    string
	CreateTime string
	UpdateTime string
	IsDeleted  string
}{
	ShowID:     "show_id",
	Channel:    "channel",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	IsDeleted:  "is_deleted",
}

// BgChannelShowApproveInfo [...]
type BgChannelShowApproveInfo struct {
	ID                       int64     `gorm:"primaryKey;column:id" json:"-"`
	CreateTime               time.Time `gorm:"column:create_time" json:"createTime"`
	IsDeleted                []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
	UpdateTime               time.Time `gorm:"column:update_time" json:"updateTime"`
	ShowApprovalNumber       string    `gorm:"column:show_approval_number" json:"showApprovalNumber"`
	ShowID                   string    `gorm:"column:show_id" json:"showId"`
	ShowName                 string    `gorm:"column:show_name" json:"showName"`
	ShowUniqueIDentification string    `gorm:"column:show_unique_identification" json:"showUniqueIdentification"`
	VenueID                  string    `gorm:"column:venue_id" json:"venueId"`
	VenueName                string    `gorm:"column:venue_name" json:"venueName"`
	CityID                   string    `gorm:"column:city_id" json:"cityId"`
	StartTime                string    `gorm:"column:start_time" json:"startTime"`       // 演出开始时间
	EndTime                  string    `gorm:"column:end_time" json:"endTime"`           // 演出结束时间
	ApprovalTime             string    `gorm:"column:approval_time" json:"approvalTime"` // 审批时间
	CityName                 string    `gorm:"column:city_name" json:"cityName"`         // 城市名称
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelShowApproveInfo) TableName() string {
	return "bg_channel_show_approve_info"
}

// BgChannelShowApproveInfoColumns get sql column name.获取数据库列名
var BgChannelShowApproveInfoColumns = struct {
	ID                       string
	CreateTime               string
	IsDeleted                string
	UpdateTime               string
	ShowApprovalNumber       string
	ShowID                   string
	ShowName                 string
	ShowUniqueIDentification string
	VenueID                  string
	VenueName                string
	CityID                   string
	StartTime                string
	EndTime                  string
	ApprovalTime             string
	CityName                 string
}{
	ID:                       "id",
	CreateTime:               "create_time",
	IsDeleted:                "is_deleted",
	UpdateTime:               "update_time",
	ShowApprovalNumber:       "show_approval_number",
	ShowID:                   "show_id",
	ShowName:                 "show_name",
	ShowUniqueIDentification: "show_unique_identification",
	VenueID:                  "venue_id",
	VenueName:                "venue_name",
	CityID:                   "city_id",
	StartTime:                "start_time",
	EndTime:                  "end_time",
	ApprovalTime:             "approval_time",
	CityName:                 "city_name",
}

// BgChannelShowOperationLog [...]
type BgChannelShowOperationLog struct {
	ID                   string    `gorm:"primaryKey;column:id" json:"-"`
	ChannelCode          string    `gorm:"column:channel_code" json:"channelCode"`
	ChannelShowMappingID string    `gorm:"column:channel_show_mapping_id" json:"channelShowMappingId"`
	InvokeID             string    `gorm:"column:invoke_id" json:"invokeId"` // 调用ID
	OpType               string    `gorm:"column:op_type" json:"opType"`     // 操作类型
	Content              string    `gorm:"column:content" json:"content"`
	ExeResult            string    `gorm:"column:exe_result" json:"exeResult"` // 执行结果
	ErrMsg               string    `gorm:"column:err_msg" json:"errMsg"`       // 错误信息
	RunStatus            string    `gorm:"column:run_status" json:"runStatus"` // 执行状态
	OperatorID           string    `gorm:"column:operator_id" json:"operatorId"`
	OperatorName         string    `gorm:"column:operator_name" json:"operatorName"`
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`
	IsDeleted            bool      `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelShowOperationLog) TableName() string {
	return "bg_channel_show_operation_log"
}

// BgChannelShowOperationLogColumns get sql column name.获取数据库列名
var BgChannelShowOperationLogColumns = struct {
	ID                   string
	ChannelCode          string
	ChannelShowMappingID string
	InvokeID             string
	OpType               string
	Content              string
	ExeResult            string
	ErrMsg               string
	RunStatus            string
	OperatorID           string
	OperatorName         string
	CreateTime           string
	IsDeleted            string
}{
	ID:                   "id",
	ChannelCode:          "channel_code",
	ChannelShowMappingID: "channel_show_mapping_id",
	InvokeID:             "invoke_id",
	OpType:               "op_type",
	Content:              "content",
	ExeResult:            "exe_result",
	ErrMsg:               "err_msg",
	RunStatus:            "run_status",
	OperatorID:           "operator_id",
	OperatorName:         "operator_name",
	CreateTime:           "create_time",
	IsDeleted:            "is_deleted",
}

// BgChannelSku [...]
type BgChannelSku struct {
	SkuID      string    `gorm:"primaryKey;column:sku_id" json:"-"`
	LeftStocks int       `gorm:"column:left_stocks" json:"leftStocks"` // 可售库存
	LockStocks int       `gorm:"column:lock_stocks" json:"lockStocks"` // 锁定库存
	SoldStocks int       `gorm:"column:sold_stocks" json:"soldStocks"` // 已售库存
	CostPrice  float64   `gorm:"column:cost_price" json:"costPrice"`   // 成本价
	Status     string    `gorm:"column:status" json:"status"`          // 状态
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted  []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelSku) TableName() string {
	return "bg_channel_sku"
}

// BgChannelSkuColumns get sql column name.获取数据库列名
var BgChannelSkuColumns = struct {
	SkuID      string
	LeftStocks string
	LockStocks string
	SoldStocks string
	CostPrice  string
	Status     string
	CreateTime string
	UpdateTime string
	IsDeleted  string
}{
	SkuID:      "sku_id",
	LeftStocks: "left_stocks",
	LockStocks: "lock_stocks",
	SoldStocks: "sold_stocks",
	CostPrice:  "cost_price",
	Status:     "status",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	IsDeleted:  "is_deleted",
}

// BgChannelSmsConfig [...]
type BgChannelSmsConfig struct {
	ID        string  `gorm:"primaryKey;column:id" json:"-"`
	ChannelID string  `gorm:"column:channelId" json:"channelId"`
	SmsKey    string  `gorm:"column:sms_key" json:"smsKey"`
	Channel   string  `gorm:"column:channel" json:"channel"`
	Platform  string  `gorm:"column:platform" json:"platform"` // 渠道平台名称
	IsSendMsg []uint8 `gorm:"column:is_send_msg" json:"isSendMsg"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelSmsConfig) TableName() string {
	return "bg_channel_sms_config"
}

// BgChannelSmsConfigColumns get sql column name.获取数据库列名
var BgChannelSmsConfigColumns = struct {
	ID        string
	ChannelID string
	SmsKey    string
	Channel   string
	Platform  string
	IsSendMsg string
}{
	ID:        "id",
	ChannelID: "channelId",
	SmsKey:    "sms_key",
	Channel:   "channel",
	Platform:  "platform",
	IsSendMsg: "is_send_msg",
}

// BgChannelSmsTemplate [...]
type BgChannelSmsTemplate struct {
	TemplateOid   int     `gorm:"primaryKey;column:template_oid" json:"-"`
	TemplateKey   string  `gorm:"column:templateKey" json:"templateKey"`
	SmsPlatform   string  `gorm:"column:smsPlatform" json:"smsPlatform"`
	TemplateValue string  `gorm:"column:templateValue" json:"templateValue"`
	Enabled       []uint8 `gorm:"column:enabled" json:"enabled"`
	Channel       string  `gorm:"column:channel" json:"channel"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelSmsTemplate) TableName() string {
	return "bg_channel_sms_template"
}

// BgChannelSmsTemplateColumns get sql column name.获取数据库列名
var BgChannelSmsTemplateColumns = struct {
	TemplateOid   string
	TemplateKey   string
	SmsPlatform   string
	TemplateValue string
	Enabled       string
	Channel       string
}{
	TemplateOid:   "template_oid",
	TemplateKey:   "templateKey",
	SmsPlatform:   "smsPlatform",
	TemplateValue: "templateValue",
	Enabled:       "enabled",
	Channel:       "channel",
}

// BgChannelStandardLocation [...]
type BgChannelStandardLocation struct {
	CityID       string    `gorm:"primaryKey;column:city_id" json:"-"`
	CityName     string    `gorm:"column:city_name" json:"cityName"`
	ProvinceID   string    `gorm:"column:province_id" json:"provinceId"`
	ProvinceName string    `gorm:"column:province_name" json:"provinceName"`
	Source       string    `gorm:"column:source" json:"source"`
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted    []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
	Version      int       `gorm:"column:version" json:"version"`
	TkingCityID  string    `gorm:"column:tking_city_id" json:"tkingCityId"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardLocation) TableName() string {
	return "bg_channel_standard_location"
}

// BgChannelStandardLocationColumns get sql column name.获取数据库列名
var BgChannelStandardLocationColumns = struct {
	CityID       string
	CityName     string
	ProvinceID   string
	ProvinceName string
	Source       string
	CreateTime   string
	UpdateTime   string
	IsDeleted    string
	Version      string
	TkingCityID  string
}{
	CityID:       "city_id",
	CityName:     "city_name",
	ProvinceID:   "province_id",
	ProvinceName: "province_name",
	Source:       "source",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
	IsDeleted:    "is_deleted",
	Version:      "version",
	TkingCityID:  "tking_city_id",
}

// BgChannelStandardOrderMappings [...]
type BgChannelStandardOrderMappings struct {
	ChannelMappingID string    `gorm:"primaryKey;column:channel_mapping_id" json:"-"` // 映射Id
	ChannelOrderID   string    `gorm:"column:channel_order_id" json:"channelOrderId"`
	MtcOrderID       string    `gorm:"column:mtc_order_id" json:"mtcOrderId"`
	ChannelCode      string    `gorm:"column:channel_code" json:"channelCode"`     // 渠道编码
	DataDirection    string    `gorm:"column:data_direction" json:"dataDirection"` // 数据流向，IN = 导入平台；OUT = 推送到第三方系统
	Status           string    `gorm:"column:status" json:"status"`                // 数据状态 更新同步失败;新增同步失败;已同步;同步中;更新待同步;更新准备中;新增待同步;新增准备中
	RetryTimes       int       `gorm:"column:retry_times" json:"retryTimes"`       // 重试次数
	ErrorCode        string    `gorm:"column:error_code" json:"errorCode"`         // 同步失败错误编码(最后一次)
	ErrorMsg         string    `gorm:"column:error_msg" json:"errorMsg"`           // 同步失败错误描述(最后一次)
	Extension        string    `gorm:"column:extension" json:"extension"`          // 扩展属性，json
	OrderNumber      string    `gorm:"column:order_number" json:"orderNumber"`     // 演出名称
	OrderStatus      string    `gorm:"column:order_status" json:"orderStatus"`     // 订单状态
	Total            float64   `gorm:"column:total" json:"total"`                  // 订单总价
	DelieveFree      float64   `gorm:"column:delieve_free" json:"delieveFree"`     // 配送费用
	UserName         string    `gorm:"column:user_name" json:"userName"`           // 收件人信息
	UserPhone        string    `gorm:"column:user_phone" json:"userPhone"`         // 用户手机号
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted        []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardOrderMappings) TableName() string {
	return "bg_channel_standard_order_mappings"
}

// BgChannelStandardOrderMappingsColumns get sql column name.获取数据库列名
var BgChannelStandardOrderMappingsColumns = struct {
	ChannelMappingID string
	ChannelOrderID   string
	MtcOrderID       string
	ChannelCode      string
	DataDirection    string
	Status           string
	RetryTimes       string
	ErrorCode        string
	ErrorMsg         string
	Extension        string
	OrderNumber      string
	OrderStatus      string
	Total            string
	DelieveFree      string
	UserName         string
	UserPhone        string
	CreateTime       string
	UpdateTime       string
	IsDeleted        string
}{
	ChannelMappingID: "channel_mapping_id",
	ChannelOrderID:   "channel_order_id",
	MtcOrderID:       "mtc_order_id",
	ChannelCode:      "channel_code",
	DataDirection:    "data_direction",
	Status:           "status",
	RetryTimes:       "retry_times",
	ErrorCode:        "error_code",
	ErrorMsg:         "error_msg",
	Extension:        "extension",
	OrderNumber:      "order_number",
	OrderStatus:      "order_status",
	Total:            "total",
	DelieveFree:      "delieve_free",
	UserName:         "user_name",
	UserPhone:        "user_phone",
	CreateTime:       "create_time",
	UpdateTime:       "update_time",
	IsDeleted:        "is_deleted",
}

// BgChannelStandardRegionMappings [...]
type BgChannelStandardRegionMappings struct {
	ChannelMappingID             string    `gorm:"primaryKey;column:channel_mapping_id" json:"-"` // 映射Id
	ChannelZoneConcreteMappingID string    `gorm:"column:channel_zone_concrete_mapping_id" json:"channelZoneConcreteMappingId"`
	MtcZoneConcreteID            string    `gorm:"column:mtc_zone_concrete_id" json:"mtcZoneConcreteId"`
	MtcSectorConcreteID          string    `gorm:"column:mtc_sector_concrete_id" json:"mtcSectorConcreteId"`
	ChannelCode                  string    `gorm:"column:channel_code" json:"channelCode"`                           // 渠道编码
	ChannelShowMappingID         string    `gorm:"column:channel_show_mapping_id" json:"channelShowMappingId"`       // 演出映射Id
	ChannelSessionMappingID      string    `gorm:"column:channel_session_mapping_id" json:"channelSessionMappingId"` // 场次映射Id
	CreateTime                   time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime                   time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted                    int8      `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardRegionMappings) TableName() string {
	return "bg_channel_standard_region_mappings"
}

// BgChannelStandardRegionMappingsColumns get sql column name.获取数据库列名
var BgChannelStandardRegionMappingsColumns = struct {
	ChannelMappingID             string
	ChannelZoneConcreteMappingID string
	MtcZoneConcreteID            string
	MtcSectorConcreteID          string
	ChannelCode                  string
	ChannelShowMappingID         string
	ChannelSessionMappingID      string
	CreateTime                   string
	UpdateTime                   string
	IsDeleted                    string
}{
	ChannelMappingID:             "channel_mapping_id",
	ChannelZoneConcreteMappingID: "channel_zone_concrete_mapping_id",
	MtcZoneConcreteID:            "mtc_zone_concrete_id",
	MtcSectorConcreteID:          "mtc_sector_concrete_id",
	ChannelCode:                  "channel_code",
	ChannelShowMappingID:         "channel_show_mapping_id",
	ChannelSessionMappingID:      "channel_session_mapping_id",
	CreateTime:                   "create_time",
	UpdateTime:                   "update_time",
	IsDeleted:                    "is_deleted",
}

// BgChannelStandardSeatPlanMappings [...]
type BgChannelStandardSeatPlanMappings struct {
	ChannelMappingID        string    `gorm:"primaryKey;column:channel_mapping_id" json:"-"` // 映射Id
	ChannelSeatPalnID       string    `gorm:"column:channel_seat_paln_id" json:"channelSeatPalnId"`
	MtcSeatPlanID           string    `gorm:"column:mtc_seat_plan_id" json:"mtcSeatPlanId"`
	ChannelCode             string    `gorm:"column:channel_code" json:"channelCode"`                           // 渠道编码
	DataDirection           string    `gorm:"column:data_direction" json:"dataDirection"`                       // 数据流向，IN = 导入平台；OUT = 推送到第三方系统
	Status                  string    `gorm:"column:status" json:"status"`                                      // 数据状态
	Extension               string    `gorm:"column:extension" json:"extension"`                                // 扩展属性，json
	OriginalPrice           float64   `gorm:"column:original_price" json:"originalPrice"`                       // 票面价
	ChannelShowMappingID    string    `gorm:"column:channel_show_mapping_id" json:"channelShowMappingId"`       // 演出场馆映射Id
	ChannelSessionMappingID string    `gorm:"column:channel_session_mapping_id" json:"channelSessionMappingId"` // 海报url
	SeatPlanType            string    `gorm:"column:seat_plan_type" json:"seatPlanType"`                        // 票介质类型
	SeatPlanUnit            string    `gorm:"column:seat_plan_unit" json:"seatPlanUnit"`                        // 票面单位
	UnitQty                 int       `gorm:"column:unit_qty" json:"unitQty"`                                   // 票面单位数量
	Limitation              int       `gorm:"column:limitation" json:"limitation"`                              // 限购数量
	BizSeatPlanID           string    `gorm:"column:biz_seat_plan_id" json:"bizSeatPlanId"`                     // 端票面id
	BizCode                 string    `gorm:"column:biz_code" json:"bizCode"`                                   // 端名称
	Comments                string    `gorm:"column:comments" json:"comments"`                                  // 票面描述
	SeatPlanName            string    `gorm:"column:seat_plan_name" json:"seatPlanName"`                        // 自定义票面显示名称
	SeatPlanStatus          string    `gorm:"column:seat_plan_status" json:"seatPlanStatus"`                    // 票面状态
	CreateTime              time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime              time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted               []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardSeatPlanMappings) TableName() string {
	return "bg_channel_standard_seat_plan_mappings"
}

// BgChannelStandardSeatPlanMappingsColumns get sql column name.获取数据库列名
var BgChannelStandardSeatPlanMappingsColumns = struct {
	ChannelMappingID        string
	ChannelSeatPalnID       string
	MtcSeatPlanID           string
	ChannelCode             string
	DataDirection           string
	Status                  string
	Extension               string
	OriginalPrice           string
	ChannelShowMappingID    string
	ChannelSessionMappingID string
	SeatPlanType            string
	SeatPlanUnit            string
	UnitQty                 string
	Limitation              string
	BizSeatPlanID           string
	BizCode                 string
	Comments                string
	SeatPlanName            string
	SeatPlanStatus          string
	CreateTime              string
	UpdateTime              string
	IsDeleted               string
}{
	ChannelMappingID:        "channel_mapping_id",
	ChannelSeatPalnID:       "channel_seat_paln_id",
	MtcSeatPlanID:           "mtc_seat_plan_id",
	ChannelCode:             "channel_code",
	DataDirection:           "data_direction",
	Status:                  "status",
	Extension:               "extension",
	OriginalPrice:           "original_price",
	ChannelShowMappingID:    "channel_show_mapping_id",
	ChannelSessionMappingID: "channel_session_mapping_id",
	SeatPlanType:            "seat_plan_type",
	SeatPlanUnit:            "seat_plan_unit",
	UnitQty:                 "unit_qty",
	Limitation:              "limitation",
	BizSeatPlanID:           "biz_seat_plan_id",
	BizCode:                 "biz_code",
	Comments:                "comments",
	SeatPlanName:            "seat_plan_name",
	SeatPlanStatus:          "seat_plan_status",
	CreateTime:              "create_time",
	UpdateTime:              "update_time",
	IsDeleted:               "is_deleted",
}

// BgChannelStandardSeatplan [...]
type BgChannelStandardSeatplan struct {
	SeatPlanID    string    `gorm:"primaryKey;column:seat_plan_id" json:"-"`
	ShowID        string    `gorm:"column:show_id" json:"showId"`
	SessionID     string    `gorm:"column:session_id" json:"sessionId"`
	OriginalPrice float64   `gorm:"column:original_price" json:"originalPrice"`
	Comment       string    `gorm:"column:comment" json:"comment"`
	UnitType      string    `gorm:"column:unit_type" json:"unitType"`
	UnitQty       int       `gorm:"column:unit_qty" json:"unitQty"` // 套票张数
	SeatPlanType  string    `gorm:"column:seat_plan_type" json:"seatPlanType"`
	Source        string    `gorm:"column:source" json:"source"`
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted     []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
	Version       int       `gorm:"column:version" json:"version"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardSeatplan) TableName() string {
	return "bg_channel_standard_seatplan"
}

// BgChannelStandardSeatplanColumns get sql column name.获取数据库列名
var BgChannelStandardSeatplanColumns = struct {
	SeatPlanID    string
	ShowID        string
	SessionID     string
	OriginalPrice string
	Comment       string
	UnitType      string
	UnitQty       string
	SeatPlanType  string
	Source        string
	CreateTime    string
	UpdateTime    string
	IsDeleted     string
	Version       string
}{
	SeatPlanID:    "seat_plan_id",
	ShowID:        "show_id",
	SessionID:     "session_id",
	OriginalPrice: "original_price",
	Comment:       "comment",
	UnitType:      "unit_type",
	UnitQty:       "unit_qty",
	SeatPlanType:  "seat_plan_type",
	Source:        "source",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
	IsDeleted:     "is_deleted",
	Version:       "version",
}

// BgChannelStandardSession [...]
type BgChannelStandardSession struct {
	SessionID        string    `gorm:"primaryKey;column:session_id" json:"-"`
	SessionBeginTime time.Time `gorm:"column:session_begin_time" json:"sessionBeginTime"` // 场次开始时间
	SessionEndTime   time.Time `gorm:"column:session_end_time" json:"sessionEndTime"`     // 场次结束时间
	ShowID           string    `gorm:"column:show_id" json:"showId"`
	SessionName      string    `gorm:"column:session_name" json:"sessionName"`
	IsPermanent      []uint8   `gorm:"column:is_permanent" json:"isPermanent"` // 是否通票
	Limitation       int       `gorm:"column:limitation" json:"limitation"`    // 限购数量
	Source           string    `gorm:"column:source" json:"source"`
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted        []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
	Version          int       `gorm:"column:version" json:"version"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardSession) TableName() string {
	return "bg_channel_standard_session"
}

// BgChannelStandardSessionColumns get sql column name.获取数据库列名
var BgChannelStandardSessionColumns = struct {
	SessionID        string
	SessionBeginTime string
	SessionEndTime   string
	ShowID           string
	SessionName      string
	IsPermanent      string
	Limitation       string
	Source           string
	CreateTime       string
	UpdateTime       string
	IsDeleted        string
	Version          string
}{
	SessionID:        "session_id",
	SessionBeginTime: "session_begin_time",
	SessionEndTime:   "session_end_time",
	ShowID:           "show_id",
	SessionName:      "session_name",
	IsPermanent:      "is_permanent",
	Limitation:       "limitation",
	Source:           "source",
	CreateTime:       "create_time",
	UpdateTime:       "update_time",
	IsDeleted:        "is_deleted",
	Version:          "version",
}

// BgChannelStandardSessionMappings [...]
type BgChannelStandardSessionMappings struct {
	ChannelMappingID     string    `gorm:"primaryKey;column:channel_mapping_id" json:"-"` // 映射Id
	ChannelSessionID     string    `gorm:"column:channel_session_id" json:"channelSessionId"`
	MtcSessionID         string    `gorm:"column:mtc_session_id" json:"mtcSessionId"`
	ChannelCode          string    `gorm:"column:channel_code" json:"channelCode"`                     // 渠道编码
	DataDirection        string    `gorm:"column:data_direction" json:"dataDirection"`                 // 数据流向，IN = 导入平台；OUT = 推送到第三方系统
	Status               string    `gorm:"column:status" json:"status"`                                // 数据状态
	Extension            string    `gorm:"column:extension" json:"extension"`                          // 扩展属性，json
	SessionStatus        string    `gorm:"column:session_status" json:"sessionStatus"`                 // 场次状态
	SessionName          string    `gorm:"column:session_name" json:"sessionName"`                     // 场次名称
	ChannelShowMappingID string    `gorm:"column:channel_show_mapping_id" json:"channelShowMappingId"` // 演出映射Id
	Limitation           int       `gorm:"column:limitation" json:"limitation"`                        // 限购数量
	BizSessionID         string    `gorm:"column:biz_session_id" json:"bizSessionId"`                  // 端场次id
	BizCode              string    `gorm:"column:biz_code" json:"bizCode"`                             // 端名称
	SessionBeginTime     time.Time `gorm:"column:session_begin_time" json:"sessionBeginTime"`          // 场次开始时间
	SessionEndTime       time.Time `gorm:"column:session_end_time" json:"sessionEndTime"`              // 场次结束时间
	IsPermanent          []uint8   `gorm:"column:is_permanent" json:"isPermanent"`                     // 是否通票
	IsCreateTemplate     []uint8   `gorm:"column:is_create_template" json:"isCreateTemplate"`          // 是否创建场次模板
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted            []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardSessionMappings) TableName() string {
	return "bg_channel_standard_session_mappings"
}

// BgChannelStandardSessionMappingsColumns get sql column name.获取数据库列名
var BgChannelStandardSessionMappingsColumns = struct {
	ChannelMappingID     string
	ChannelSessionID     string
	MtcSessionID         string
	ChannelCode          string
	DataDirection        string
	Status               string
	Extension            string
	SessionStatus        string
	SessionName          string
	ChannelShowMappingID string
	Limitation           string
	BizSessionID         string
	BizCode              string
	SessionBeginTime     string
	SessionEndTime       string
	IsPermanent          string
	IsCreateTemplate     string
	CreateTime           string
	UpdateTime           string
	IsDeleted            string
}{
	ChannelMappingID:     "channel_mapping_id",
	ChannelSessionID:     "channel_session_id",
	MtcSessionID:         "mtc_session_id",
	ChannelCode:          "channel_code",
	DataDirection:        "data_direction",
	Status:               "status",
	Extension:            "extension",
	SessionStatus:        "session_status",
	SessionName:          "session_name",
	ChannelShowMappingID: "channel_show_mapping_id",
	Limitation:           "limitation",
	BizSessionID:         "biz_session_id",
	BizCode:              "biz_code",
	SessionBeginTime:     "session_begin_time",
	SessionEndTime:       "session_end_time",
	IsPermanent:          "is_permanent",
	IsCreateTemplate:     "is_create_template",
	CreateTime:           "create_time",
	UpdateTime:           "update_time",
	IsDeleted:            "is_deleted",
}

// BgChannelStandardShow [...]
type BgChannelStandardShow struct {
	ShowID               string    `gorm:"primaryKey;column:show_id" json:"-"`
	ShowName             string    `gorm:"column:show_name" json:"showName"`
	ShowType             string    `gorm:"column:show_type" json:"showType"`
	VenueID              string    `gorm:"column:venue_id" json:"venueId"`
	VenueName            string    `gorm:"column:venue_name" json:"venueName"`
	FirstShowTime        string    `gorm:"column:first_show_time" json:"firstShowTime"`
	LastShowTime         string    `gorm:"column:last_show_time" json:"lastShowTime"`
	Content              string    `gorm:"column:content" json:"content"`
	Description          string    `gorm:"column:description" json:"description"`
	SupportSeatPicking   []uint8   `gorm:"column:support_seat_picking" json:"supportSeatPicking"`
	PosterURL            string    `gorm:"column:poster_url" json:"posterUrl"`
	CityID               string    `gorm:"column:city_id" json:"cityId"`
	CityName             string    `gorm:"column:city_name" json:"cityName"`
	Source               string    `gorm:"column:source" json:"source"`
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted            []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
	IDentificationCode   string    `gorm:"column:identification_code" json:"identificationCode"`      // 文广局、公安、执法总队批准演出销售后会下发的识别码
	SeatPlanURL          string    `gorm:"column:seat_plan_url" json:"seatPlanUrl"`                   // 演出场馆图url
	IDentityRequiredType int       `gorm:"column:identity_required_type" json:"identityRequiredType"` // 购买此演出身份证使用类型 0 不需要使用身份证 1 多张票仅需一张身份证 2 每张票对应一张身份证
	SupportDeliverMethod int       `gorm:"column:support_deliver_method" json:"supportDeliverMethod"` // 支持的配送方式 按2进制位表示 由低到高位分别是快递 上门 现场
	ApprovalNo           string    `gorm:"column:approval_no" json:"approvalNo"`                      // 演出批文号
	ApprovalPhotoURL     string    `gorm:"column:approval_photo_url" json:"approvalPhotoUrl"`         // 批文图url
	OfficialNoticeURL    string    `gorm:"column:official_notice_url" json:"officialNoticeUrl"`       // 主办公告文件URL
	Version              int       `gorm:"column:version" json:"version"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardShow) TableName() string {
	return "bg_channel_standard_show"
}

// BgChannelStandardShowColumns get sql column name.获取数据库列名
var BgChannelStandardShowColumns = struct {
	ShowID               string
	ShowName             string
	ShowType             string
	VenueID              string
	VenueName            string
	FirstShowTime        string
	LastShowTime         string
	Content              string
	Description          string
	SupportSeatPicking   string
	PosterURL            string
	CityID               string
	CityName             string
	Source               string
	CreateTime           string
	UpdateTime           string
	IsDeleted            string
	IDentificationCode   string
	SeatPlanURL          string
	IDentityRequiredType string
	SupportDeliverMethod string
	ApprovalNo           string
	ApprovalPhotoURL     string
	OfficialNoticeURL    string
	Version              string
}{
	ShowID:               "show_id",
	ShowName:             "show_name",
	ShowType:             "show_type",
	VenueID:              "venue_id",
	VenueName:            "venue_name",
	FirstShowTime:        "first_show_time",
	LastShowTime:         "last_show_time",
	Content:              "content",
	Description:          "description",
	SupportSeatPicking:   "support_seat_picking",
	PosterURL:            "poster_url",
	CityID:               "city_id",
	CityName:             "city_name",
	Source:               "source",
	CreateTime:           "create_time",
	UpdateTime:           "update_time",
	IsDeleted:            "is_deleted",
	IDentificationCode:   "identification_code",
	SeatPlanURL:          "seat_plan_url",
	IDentityRequiredType: "identity_required_type",
	SupportDeliverMethod: "support_deliver_method",
	ApprovalNo:           "approval_no",
	ApprovalPhotoURL:     "approval_photo_url",
	OfficialNoticeURL:    "official_notice_url",
	Version:              "version",
}

// BgChannelStandardShowMappings [...]
type BgChannelStandardShowMappings struct {
	ChannelMappingID     string    `gorm:"primaryKey;column:channel_mapping_id" json:"-"` // 映射Id
	ChannelShowID        string    `gorm:"column:channel_show_id" json:"channelShowId"`
	MtcShowID            string    `gorm:"column:mtc_show_id" json:"mtcShowId"`                       // MTC_SHOW_ID
	ChannelCode          string    `gorm:"column:channel_code" json:"channelCode"`                    // 渠道编码
	DataDirection        string    `gorm:"column:data_direction" json:"dataDirection"`                // 数据流向，IN = 导入平台；OUT = 推送到第三方系统
	Status               string    `gorm:"column:status" json:"status"`                               // 数据状态 更新同步失败;新增同步失败;已同步;同步中;更新待同步;更新准备中;新增待同步;新增准备中
	RetryTimes           int       `gorm:"column:retry_times" json:"retryTimes"`                      // 重试次数
	ErrorCode            string    `gorm:"column:error_code" json:"errorCode"`                        // 同步失败错误编码(最后一次)
	ErrorMsg             string    `gorm:"column:error_msg" json:"errorMsg"`                          // 同步失败错误描述(最后一次)
	Extension            string    `gorm:"column:extension" json:"extension"`                         // 扩展属性，json
	ShowName             string    `gorm:"column:show_name" json:"showName"`                          // 演出名称
	ShowVenueMappingID   string    `gorm:"column:show_venue_mapping_id" json:"showVenueMappingId"`    // 演出场馆映射Id
	PosterURL            string    `gorm:"column:poster_url" json:"posterUrl"`                        // 海报url
	PosterURL1           string    `gorm:"column:poster_url1" json:"posterUrl1"`                      // 海报补充
	Content              string    `gorm:"column:content" json:"content"`                             // 演出详情
	SeatPlanURL          string    `gorm:"column:seat_plan_url" json:"seatPlanUrl"`                   // 演出场馆图url
	Limitation           int       `gorm:"column:limitation" json:"limitation"`                       // 限购数量
	BizShowID            string    `gorm:"column:biz_show_id" json:"bizShowId"`                       // 端演出id
	BizCode              string    `gorm:"column:biz_code" json:"bizCode"`                            // 端名称
	ShowType             string    `gorm:"column:show_type" json:"showType"`                          // 演出类型
	IDentityRequiredType int       `gorm:"column:identity_required_type" json:"identityRequiredType"` // 购买此演出身份证使用类型 0 不需要使用身份证 1 多张票仅需一张身份证 2 每张票对应一张身份证
	SupportDeliverMethod int       `gorm:"column:support_deliver_method" json:"supportDeliverMethod"` // 支持的配送方式 按2进制位表示 由低到高位分别是快递 上门 现场
	ApprovalNo           string    `gorm:"column:approval_no" json:"approvalNo"`                      // 演出批文号
	ApprovalPhotoURL     string    `gorm:"column:approval_photo_url" json:"approvalPhotoUrl"`         // 批文图url
	OfficialNoticeURL    string    `gorm:"column:official_notice_url" json:"officialNoticeUrl"`       // 主办公告文件URL
	IDentificationCode   string    `gorm:"column:identification_code" json:"identificationCode"`      // 文广局、公安、执法总队批准演出销售后会下发的识别码
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted            []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
	SeatPickType         string    `gorm:"column:seat_pick_type" json:"seatPickType"` // 选座标识（默认不支持选座）
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardShowMappings) TableName() string {
	return "bg_channel_standard_show_mappings"
}

// BgChannelStandardShowMappingsColumns get sql column name.获取数据库列名
var BgChannelStandardShowMappingsColumns = struct {
	ChannelMappingID     string
	ChannelShowID        string
	MtcShowID            string
	ChannelCode          string
	DataDirection        string
	Status               string
	RetryTimes           string
	ErrorCode            string
	ErrorMsg             string
	Extension            string
	ShowName             string
	ShowVenueMappingID   string
	PosterURL            string
	PosterURL1           string
	Content              string
	SeatPlanURL          string
	Limitation           string
	BizShowID            string
	BizCode              string
	ShowType             string
	IDentityRequiredType string
	SupportDeliverMethod string
	ApprovalNo           string
	ApprovalPhotoURL     string
	OfficialNoticeURL    string
	IDentificationCode   string
	CreateTime           string
	UpdateTime           string
	IsDeleted            string
	SeatPickType         string
}{
	ChannelMappingID:     "channel_mapping_id",
	ChannelShowID:        "channel_show_id",
	MtcShowID:            "mtc_show_id",
	ChannelCode:          "channel_code",
	DataDirection:        "data_direction",
	Status:               "status",
	RetryTimes:           "retry_times",
	ErrorCode:            "error_code",
	ErrorMsg:             "error_msg",
	Extension:            "extension",
	ShowName:             "show_name",
	ShowVenueMappingID:   "show_venue_mapping_id",
	PosterURL:            "poster_url",
	PosterURL1:           "poster_url1",
	Content:              "content",
	SeatPlanURL:          "seat_plan_url",
	Limitation:           "limitation",
	BizShowID:            "biz_show_id",
	BizCode:              "biz_code",
	ShowType:             "show_type",
	IDentityRequiredType: "identity_required_type",
	SupportDeliverMethod: "support_deliver_method",
	ApprovalNo:           "approval_no",
	ApprovalPhotoURL:     "approval_photo_url",
	OfficialNoticeURL:    "official_notice_url",
	IDentificationCode:   "identification_code",
	CreateTime:           "create_time",
	UpdateTime:           "update_time",
	IsDeleted:            "is_deleted",
	SeatPickType:         "seat_pick_type",
}

// BgChannelStandardTicket [...]
type BgChannelStandardTicket struct {
	TicketID           string    `gorm:"primaryKey;column:ticket_id" json:"-"`
	SeatPlanID         string    `gorm:"column:seat_plan_id" json:"seatPlanId"`
	SessionID          string    `gorm:"column:session_id" json:"sessionId"`
	ShowID             string    `gorm:"column:show_id" json:"showId"`
	TicketType         string    `gorm:"column:ticket_type" json:"ticketType"` // 纸质票/电子票
	TicketStatus       string    `gorm:"column:ticket_status" json:"ticketStatus"`
	Price              float64   `gorm:"column:price" json:"price"`               // 报价
	LeftStock          int       `gorm:"column:left_stock" json:"leftStock"`      // 可售库存
	Row                int       `gorm:"column:row" json:"row"`                   // 排号
	StartSeatNo        int       `gorm:"column:start_seat_no" json:"startSeatNo"` // 开始座位号
	EndSeatNo          int       `gorm:"column:end_seat_no" json:"endSeatNo"`     // 座位结束号
	SupportSeatPicking []uint8   `gorm:"column:support_seat_picking" json:"supportSeatPicking"`
	Source             string    `gorm:"column:source" json:"source"`
	CreateTime         time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime         time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted          []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
	Version            int       `gorm:"column:version" json:"version"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardTicket) TableName() string {
	return "bg_channel_standard_ticket"
}

// BgChannelStandardTicketColumns get sql column name.获取数据库列名
var BgChannelStandardTicketColumns = struct {
	TicketID           string
	SeatPlanID         string
	SessionID          string
	ShowID             string
	TicketType         string
	TicketStatus       string
	Price              string
	LeftStock          string
	Row                string
	StartSeatNo        string
	EndSeatNo          string
	SupportSeatPicking string
	Source             string
	CreateTime         string
	UpdateTime         string
	IsDeleted          string
	Version            string
}{
	TicketID:           "ticket_id",
	SeatPlanID:         "seat_plan_id",
	SessionID:          "session_id",
	ShowID:             "show_id",
	TicketType:         "ticket_type",
	TicketStatus:       "ticket_status",
	Price:              "price",
	LeftStock:          "left_stock",
	Row:                "row",
	StartSeatNo:        "start_seat_no",
	EndSeatNo:          "end_seat_no",
	SupportSeatPicking: "support_seat_picking",
	Source:             "source",
	CreateTime:         "create_time",
	UpdateTime:         "update_time",
	IsDeleted:          "is_deleted",
	Version:            "version",
}

// BgChannelStandardTicketMappings [...]
type BgChannelStandardTicketMappings struct {
	ChannelMappingID         string    `gorm:"primaryKey;column:channel_mapping_id" json:"-"` // 映射Id
	ChannelTicketID          string    `gorm:"column:channel_ticket_id" json:"channelTicketId"`
	MtcTicketID              string    `gorm:"column:mtc_ticket_id" json:"mtcTicketId"`                             // MTC_SHOW_ID
	ChannelCode              string    `gorm:"column:channel_code" json:"channelCode"`                              // 渠道编码
	DataDirection            string    `gorm:"column:data_direction" json:"dataDirection"`                          // 数据流向，IN = 导入平台；OUT = 推送到第三方系统
	Status                   string    `gorm:"column:status" json:"status"`                                         // 数据状态 更新同步失败;新增同步失败;已同步;同步中;更新待同步;更新准备中;新增待同步;新增准备中
	RetryTimes               int       `gorm:"column:retry_times" json:"retryTimes"`                                // 重试次数
	ErrorCode                string    `gorm:"column:error_code" json:"errorCode"`                                  // 同步失败错误编码(最后一次)
	ErrorMsg                 string    `gorm:"column:error_msg" json:"errorMsg"`                                    // 同步失败错误描述(最后一次)
	Extension                string    `gorm:"column:extension" json:"extension"`                                   // 扩展属性，json
	ChannelShowMappingID     string    `gorm:"column:channel_show_mapping_id" json:"channelShowMappingId"`          // 演出映射Id
	ChannelSessionMappingID  string    `gorm:"column:channel_session_mapping_id" json:"channelSessionMappingId"`    // 场次映射Id
	ChannelSeatPlanMappingID string    `gorm:"column:channel_seat_plan_mapping_id" json:"channelSeatPlanMappingId"` // 票面映射Id
	TicketType               string    `gorm:"column:ticket_type" json:"ticketType"`                                // 纸质票/电子票
	TicketStatus             string    `gorm:"column:ticket_status" json:"ticketStatus"`
	Price                    float64   `gorm:"column:price" json:"price"`                   // 报价
	SalePrice                float64   `gorm:"column:sale_price" json:"salePrice"`          // 售价
	LeftStock                int       `gorm:"column:left_stock" json:"leftStock"`          // 可售库存
	IsSupportSeat            []uint8   `gorm:"column:is_support_seat" json:"isSupportSeat"` // 是否支持选座
	Row                      string    `gorm:"column:row" json:"row"`                       // 排号
	StartSeatNo              string    `gorm:"column:start_seat_no" json:"startSeatNo"`     // 开始座位号
	EndSeatNo                string    `gorm:"column:end_seat_no" json:"endSeatNo"`         // 座位结束号
	CreateTime               time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime               time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted                []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardTicketMappings) TableName() string {
	return "bg_channel_standard_ticket_mappings"
}

// BgChannelStandardTicketMappingsColumns get sql column name.获取数据库列名
var BgChannelStandardTicketMappingsColumns = struct {
	ChannelMappingID         string
	ChannelTicketID          string
	MtcTicketID              string
	ChannelCode              string
	DataDirection            string
	Status                   string
	RetryTimes               string
	ErrorCode                string
	ErrorMsg                 string
	Extension                string
	ChannelShowMappingID     string
	ChannelSessionMappingID  string
	ChannelSeatPlanMappingID string
	TicketType               string
	TicketStatus             string
	Price                    string
	SalePrice                string
	LeftStock                string
	IsSupportSeat            string
	Row                      string
	StartSeatNo              string
	EndSeatNo                string
	CreateTime               string
	UpdateTime               string
	IsDeleted                string
}{
	ChannelMappingID:         "channel_mapping_id",
	ChannelTicketID:          "channel_ticket_id",
	MtcTicketID:              "mtc_ticket_id",
	ChannelCode:              "channel_code",
	DataDirection:            "data_direction",
	Status:                   "status",
	RetryTimes:               "retry_times",
	ErrorCode:                "error_code",
	ErrorMsg:                 "error_msg",
	Extension:                "extension",
	ChannelShowMappingID:     "channel_show_mapping_id",
	ChannelSessionMappingID:  "channel_session_mapping_id",
	ChannelSeatPlanMappingID: "channel_seat_plan_mapping_id",
	TicketType:               "ticket_type",
	TicketStatus:             "ticket_status",
	Price:                    "price",
	SalePrice:                "sale_price",
	LeftStock:                "left_stock",
	IsSupportSeat:            "is_support_seat",
	Row:                      "row",
	StartSeatNo:              "start_seat_no",
	EndSeatNo:                "end_seat_no",
	CreateTime:               "create_time",
	UpdateTime:               "update_time",
	IsDeleted:                "is_deleted",
}

// BgChannelStandardTicketSeat [...]
type BgChannelStandardTicketSeat struct {
	SeatID     string    `gorm:"column:seat_id" json:"seatId"`
	TicketID   string    `gorm:"column:ticket_id" json:"ticketId"` // MTC_SHOW_ID
	GraphCol   int       `gorm:"column:graph_col" json:"graphCol"` // 物理列号
	GraphRow   int       `gorm:"column:graph_row" json:"graphRow"` // 物理行号
	Row        int       `gorm:"column:row" json:"row"`            // 座位排号
	Col        int       `gorm:"column:col" json:"col"`            // 座位列号
	SeatNo     string    `gorm:"column:seat_no" json:"seatNo"`     // 座位编码
	Source     string    `gorm:"column:source" json:"source"`
	Version    int       `gorm:"column:version" json:"version"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted  []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardTicketSeat) TableName() string {
	return "bg_channel_standard_ticket_seat"
}

// BgChannelStandardTicketSeatColumns get sql column name.获取数据库列名
var BgChannelStandardTicketSeatColumns = struct {
	SeatID     string
	TicketID   string
	GraphCol   string
	GraphRow   string
	Row        string
	Col        string
	SeatNo     string
	Source     string
	Version    string
	CreateTime string
	UpdateTime string
	IsDeleted  string
}{
	SeatID:     "seat_id",
	TicketID:   "ticket_id",
	GraphCol:   "graph_col",
	GraphRow:   "graph_row",
	Row:        "row",
	Col:        "col",
	SeatNo:     "seat_no",
	Source:     "source",
	Version:    "version",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	IsDeleted:  "is_deleted",
}

// BgChannelStandardTicketSeatMappings [...]
type BgChannelStandardTicketSeatMappings struct {
	ChannelMappingID             string    `gorm:"primaryKey;column:channel_mapping_id" json:"-"` // 映射Id
	ChannelSeatID                string    `gorm:"column:channel_seat_id" json:"channelSeatId"`
	MtcSeatID                    string    `gorm:"column:mtc_seat_id" json:"mtcSeatId"`
	ChannelCode                  string    `gorm:"column:channel_code" json:"channelCode"`     // 渠道编码
	DataDirection                string    `gorm:"column:data_direction" json:"dataDirection"` // 数据流向，IN = 导入平台；OUT = 推送到第三方系统
	Status                       string    `gorm:"column:status" json:"status"`
	Extension                    string    `gorm:"column:extension" json:"extension"`                                   // 扩展属性，json
	ChannelShowMappingID         string    `gorm:"column:channel_show_mapping_id" json:"channelShowMappingId"`          // 演出映射Id
	ChannelSessionMappingID      string    `gorm:"column:channel_session_mapping_id" json:"channelSessionMappingId"`    // 场次映射Id
	ChannelSeatPlanMappingID     string    `gorm:"column:channel_seat_plan_mapping_id" json:"channelSeatPlanMappingId"` // 票面映射Id
	ChannelTicketMappingID       string    `gorm:"column:channel_ticket_mapping_id" json:"channelTicketMappingId"`
	ChannelZoneConcreteMappingID string    `gorm:"column:channel_zone_concrete_mapping_id" json:"channelZoneConcreteMappingId"`
	ChannelRegionMappingID       string    `gorm:"column:channel_region_mapping_id" json:"channelRegionMappingId"`
	MtcZoneSeatConcreteID        string    `gorm:"column:mtc_zone_seat_concrete_id" json:"mtcZoneSeatConcreteId"`
	TicketStatus                 string    `gorm:"column:ticket_status" json:"ticketStatus"` // 状态
	GraphCol                     int       `gorm:"column:graph_col" json:"graphCol"`         // 物理列号
	GraphRow                     int       `gorm:"column:graph_row" json:"graphRow"`         // 物理行号
	Row                          string    `gorm:"column:row" json:"row"`                    // 座位排号
	Col                          string    `gorm:"column:col" json:"col"`                    // 座位列号
	SeatNo                       string    `gorm:"column:seat_no" json:"seatNo"`             // 座位编码
	Remark                       string    `gorm:"column:remark" json:"remark"`
	CreateTime                   time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime                   time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted                    []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardTicketSeatMappings) TableName() string {
	return "bg_channel_standard_ticket_seat_mappings"
}

// BgChannelStandardTicketSeatMappingsColumns get sql column name.获取数据库列名
var BgChannelStandardTicketSeatMappingsColumns = struct {
	ChannelMappingID             string
	ChannelSeatID                string
	MtcSeatID                    string
	ChannelCode                  string
	DataDirection                string
	Status                       string
	Extension                    string
	ChannelShowMappingID         string
	ChannelSessionMappingID      string
	ChannelSeatPlanMappingID     string
	ChannelTicketMappingID       string
	ChannelZoneConcreteMappingID string
	ChannelRegionMappingID       string
	MtcZoneSeatConcreteID        string
	TicketStatus                 string
	GraphCol                     string
	GraphRow                     string
	Row                          string
	Col                          string
	SeatNo                       string
	Remark                       string
	CreateTime                   string
	UpdateTime                   string
	IsDeleted                    string
}{
	ChannelMappingID:             "channel_mapping_id",
	ChannelSeatID:                "channel_seat_id",
	MtcSeatID:                    "mtc_seat_id",
	ChannelCode:                  "channel_code",
	DataDirection:                "data_direction",
	Status:                       "status",
	Extension:                    "extension",
	ChannelShowMappingID:         "channel_show_mapping_id",
	ChannelSessionMappingID:      "channel_session_mapping_id",
	ChannelSeatPlanMappingID:     "channel_seat_plan_mapping_id",
	ChannelTicketMappingID:       "channel_ticket_mapping_id",
	ChannelZoneConcreteMappingID: "channel_zone_concrete_mapping_id",
	ChannelRegionMappingID:       "channel_region_mapping_id",
	MtcZoneSeatConcreteID:        "mtc_zone_seat_concrete_id",
	TicketStatus:                 "ticket_status",
	GraphCol:                     "graph_col",
	GraphRow:                     "graph_row",
	Row:                          "row",
	Col:                          "col",
	SeatNo:                       "seat_no",
	Remark:                       "remark",
	CreateTime:                   "create_time",
	UpdateTime:                   "update_time",
	IsDeleted:                    "is_deleted",
}

// BgChannelStandardVenue [...]
type BgChannelStandardVenue struct {
	VenueID      string    `gorm:"primaryKey;column:venue_id" json:"-"`
	VenueName    string    `gorm:"column:venue_name" json:"venueName"`
	Address      string    `gorm:"column:address" json:"address"`
	CityID       string    `gorm:"column:city_id" json:"cityId"`
	CityName     string    `gorm:"column:city_name" json:"cityName"`
	Lng          string    `gorm:"column:lng" json:"lng"`
	Lat          string    `gorm:"column:lat" json:"lat"`
	Source       string    `gorm:"column:source" json:"source"`
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted    []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
	TkingVenueID string    `gorm:"column:tking_venue_id" json:"tkingVenueId"`
	Version      int       `gorm:"column:version" json:"version"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardVenue) TableName() string {
	return "bg_channel_standard_venue"
}

// BgChannelStandardVenueColumns get sql column name.获取数据库列名
var BgChannelStandardVenueColumns = struct {
	VenueID      string
	VenueName    string
	Address      string
	CityID       string
	CityName     string
	Lng          string
	Lat          string
	Source       string
	CreateTime   string
	UpdateTime   string
	IsDeleted    string
	TkingVenueID string
	Version      string
}{
	VenueID:      "venue_id",
	VenueName:    "venue_name",
	Address:      "address",
	CityID:       "city_id",
	CityName:     "city_name",
	Lng:          "lng",
	Lat:          "lat",
	Source:       "source",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
	IsDeleted:    "is_deleted",
	TkingVenueID: "tking_venue_id",
	Version:      "version",
}

// BgChannelStandardVenueMappings [...]
type BgChannelStandardVenueMappings struct {
	ChannelMappingID    string    `gorm:"primaryKey;column:channel_mapping_id" json:"-"` // 映射Id
	ChannelVenueID      string    `gorm:"column:channel_venue_id" json:"channelVenueId"`
	MtcVenueID          string    `gorm:"column:mtc_venue_id" json:"mtcVenueId"`                   // MTC_SHOW_ID
	ChannelCode         string    `gorm:"column:channel_code" json:"channelCode"`                  // 渠道编码
	DataDirection       string    `gorm:"column:data_direction" json:"dataDirection"`              // 数据流向，IN = 导入平台；OUT = 推送到第三方系统
	Status              string    `gorm:"column:status" json:"status"`                             // 数据状态 更新同步失败;新增同步失败;已同步;同步中;更新待同步;更新准备中;新增待同步;新增准备中
	RetryTimes          int       `gorm:"column:retry_times" json:"retryTimes"`                    // 重试次数
	ErrorCode           string    `gorm:"column:error_code" json:"errorCode"`                      // 同步失败错误编码(最后一次)
	ErrorMsg            string    `gorm:"column:error_msg" json:"errorMsg"`                        // 同步失败错误描述(最后一次)
	Extension           string    `gorm:"column:extension" json:"extension"`                       // 扩展属性，json
	VenueName           string    `gorm:"column:venue_name" json:"venueName"`                      // 场馆名称
	Lng                 string    `gorm:"column:lng" json:"lng"`                                   // 经度
	Lat                 string    `gorm:"column:lat" json:"lat"`                                   // 纬度
	Address             string    `gorm:"column:address" json:"address"`                           // 地址
	CityID              string    `gorm:"column:city_id" json:"cityId"`                            // 城市Id
	SeatLoadRule        string    `gorm:"column:seat_load_rule" json:"seatLoadRule"`               // 场馆座位加载策略 GLOBAL-全局加载 REGION-分区加载
	VenueTemplateID     string    `gorm:"column:venue_template_id" json:"venueTemplateId"`         // 场馆中心场馆模板id
	VenueTemplateStatus string    `gorm:"column:venue_template_status" json:"venueTemplateStatus"` // 场馆中心场馆模板, DONE-已创建模板，NOT_DONE-未创建
	CreateTime          time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime          time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted           []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardVenueMappings) TableName() string {
	return "bg_channel_standard_venue_mappings"
}

// BgChannelStandardVenueMappingsColumns get sql column name.获取数据库列名
var BgChannelStandardVenueMappingsColumns = struct {
	ChannelMappingID    string
	ChannelVenueID      string
	MtcVenueID          string
	ChannelCode         string
	DataDirection       string
	Status              string
	RetryTimes          string
	ErrorCode           string
	ErrorMsg            string
	Extension           string
	VenueName           string
	Lng                 string
	Lat                 string
	Address             string
	CityID              string
	SeatLoadRule        string
	VenueTemplateID     string
	VenueTemplateStatus string
	CreateTime          string
	UpdateTime          string
	IsDeleted           string
}{
	ChannelMappingID:    "channel_mapping_id",
	ChannelVenueID:      "channel_venue_id",
	MtcVenueID:          "mtc_venue_id",
	ChannelCode:         "channel_code",
	DataDirection:       "data_direction",
	Status:              "status",
	RetryTimes:          "retry_times",
	ErrorCode:           "error_code",
	ErrorMsg:            "error_msg",
	Extension:           "extension",
	VenueName:           "venue_name",
	Lng:                 "lng",
	Lat:                 "lat",
	Address:             "address",
	CityID:              "city_id",
	SeatLoadRule:        "seat_load_rule",
	VenueTemplateID:     "venue_template_id",
	VenueTemplateStatus: "venue_template_status",
	CreateTime:          "create_time",
	UpdateTime:          "update_time",
	IsDeleted:           "is_deleted",
}

// BgChannelStandardVenueMappingsCopy [...]
type BgChannelStandardVenueMappingsCopy struct {
	ChannelMappingID string    `gorm:"primaryKey;column:channel_mapping_id" json:"-"` // 映射Id
	ChannelVenueID   string    `gorm:"column:channel_venue_id" json:"channelVenueId"`
	MtcVenueID       string    `gorm:"column:mtc_venue_id" json:"mtcVenueId"`      // MTC_SHOW_ID
	ChannelCode      string    `gorm:"column:channel_code" json:"channelCode"`     // 渠道编码
	DataDirection    string    `gorm:"column:data_direction" json:"dataDirection"` // 数据流向，IN = 导入平台；OUT = 推送到第三方系统
	Status           string    `gorm:"column:status" json:"status"`                // 数据状态 更新同步失败;新增同步失败;已同步;同步中;更新待同步;更新准备中;新增待同步;新增准备中
	RetryTimes       int       `gorm:"column:retry_times" json:"retryTimes"`       // 重试次数
	ErrorCode        string    `gorm:"column:error_code" json:"errorCode"`         // 同步失败错误编码(最后一次)
	ErrorMsg         string    `gorm:"column:error_msg" json:"errorMsg"`           // 同步失败错误描述(最后一次)
	Extension        string    `gorm:"column:extension" json:"extension"`          // 扩展属性，json
	VenueName        string    `gorm:"column:venue_name" json:"venueName"`         // 场馆名称
	Lng              string    `gorm:"column:lng" json:"lng"`                      // 经度
	Lat              string    `gorm:"column:lat" json:"lat"`                      // 纬度
	Address          string    `gorm:"column:address" json:"address"`              // 地址
	CityID           string    `gorm:"column:city_id" json:"cityId"`               // 城市Id
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted        []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardVenueMappingsCopy) TableName() string {
	return "bg_channel_standard_venue_mappings_copy"
}

// BgChannelStandardVenueMappingsCopyColumns get sql column name.获取数据库列名
var BgChannelStandardVenueMappingsCopyColumns = struct {
	ChannelMappingID string
	ChannelVenueID   string
	MtcVenueID       string
	ChannelCode      string
	DataDirection    string
	Status           string
	RetryTimes       string
	ErrorCode        string
	ErrorMsg         string
	Extension        string
	VenueName        string
	Lng              string
	Lat              string
	Address          string
	CityID           string
	CreateTime       string
	UpdateTime       string
	IsDeleted        string
}{
	ChannelMappingID: "channel_mapping_id",
	ChannelVenueID:   "channel_venue_id",
	MtcVenueID:       "mtc_venue_id",
	ChannelCode:      "channel_code",
	DataDirection:    "data_direction",
	Status:           "status",
	RetryTimes:       "retry_times",
	ErrorCode:        "error_code",
	ErrorMsg:         "error_msg",
	Extension:        "extension",
	VenueName:        "venue_name",
	Lng:              "lng",
	Lat:              "lat",
	Address:          "address",
	CityID:           "city_id",
	CreateTime:       "create_time",
	UpdateTime:       "update_time",
	IsDeleted:        "is_deleted",
}

// BgChannelStandardZoneConcrete [...]
type BgChannelStandardZoneConcrete struct {
	ZoneConcreteID   string    `gorm:"column:zone_concrete_id" json:"zoneConcreteId"`
	ZoneConcreteName string    `gorm:"column:zone_concrete_name" json:"zoneConcreteName"` // 看台名称
	ZoneConcreteCode string    `gorm:"column:zone_concrete_code" json:"zoneConcreteCode"` // 看台Id
	SessionID        string    `gorm:"column:session_id" json:"sessionId"`
	ShowID           string    `gorm:"column:show_id" json:"showId"`
	FloorNo          int       `gorm:"column:floor_no" json:"floorNo"` // 楼层
	Source           string    `gorm:"column:source" json:"source"`
	Version          int       `gorm:"column:version" json:"version"`
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted        []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardZoneConcrete) TableName() string {
	return "bg_channel_standard_zone_concrete"
}

// BgChannelStandardZoneConcreteColumns get sql column name.获取数据库列名
var BgChannelStandardZoneConcreteColumns = struct {
	ZoneConcreteID   string
	ZoneConcreteName string
	ZoneConcreteCode string
	SessionID        string
	ShowID           string
	FloorNo          string
	Source           string
	Version          string
	CreateTime       string
	UpdateTime       string
	IsDeleted        string
}{
	ZoneConcreteID:   "zone_concrete_id",
	ZoneConcreteName: "zone_concrete_name",
	ZoneConcreteCode: "zone_concrete_code",
	SessionID:        "session_id",
	ShowID:           "show_id",
	FloorNo:          "floor_no",
	Source:           "source",
	Version:          "version",
	CreateTime:       "create_time",
	UpdateTime:       "update_time",
	IsDeleted:        "is_deleted",
}

// BgChannelStandardZoneConcreteMappings [...]
type BgChannelStandardZoneConcreteMappings struct {
	ChannelMappingID        string    `gorm:"primaryKey;column:channel_mapping_id" json:"-"`                // 映射Id
	ChannelRowConcreteID    string    `gorm:"column:channel_row_concrete_id" json:"channelRowConcreteId"`   // 渠道二级区域id
	ChannelZoneConcreteID   string    `gorm:"column:channel_zone_concrete_id" json:"channelZoneConcreteId"` // 渠道一级区域id
	MtcZoneConcreteID       string    `gorm:"column:mtc_zone_concrete_id" json:"mtcZoneConcreteId"`
	MtcSectorConcreteID     string    `gorm:"column:mtc_sector_concrete_id" json:"mtcSectorConcreteId"`
	ChannelCode             string    `gorm:"column:channel_code" json:"channelCode"`                           // 渠道编码
	DataDirection           string    `gorm:"column:data_direction" json:"dataDirection"`                       // 数据流向，IN = 导入平台；OUT = 推送到第三方系统
	Status                  string    `gorm:"column:status" json:"status"`                                      // 数据状态 更新同步失败;新增同步失败;已同步;同步中;更新待同步;更新准备中;新增待同步;新增准备中
	Extension               string    `gorm:"column:extension" json:"extension"`                                // 扩展属性，json
	ChannelShowMappingID    string    `gorm:"column:channel_show_mapping_id" json:"channelShowMappingId"`       // 演出映射Id
	ChannelSessionMappingID string    `gorm:"column:channel_session_mapping_id" json:"channelSessionMappingId"` // 场次映射Id
	ZoneConcreteName        string    `gorm:"column:zone_concrete_name" json:"zoneConcreteName"`                // 看台名称
	EndNo                   string    `gorm:"column:end_no" json:"endNo"`
	RowName                 string    `gorm:"column:row_name" json:"rowName"`
	ZoneConcreteCode        string    `gorm:"column:zone_concrete_code" json:"zoneConcreteCode"` // 看台Id
	FloorNo                 int       `gorm:"column:floor_no" json:"floorNo"`                    // 楼层
	CreateTime              time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime              time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted               []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardZoneConcreteMappings) TableName() string {
	return "bg_channel_standard_zone_concrete_mappings"
}

// BgChannelStandardZoneConcreteMappingsColumns get sql column name.获取数据库列名
var BgChannelStandardZoneConcreteMappingsColumns = struct {
	ChannelMappingID        string
	ChannelRowConcreteID    string
	ChannelZoneConcreteID   string
	MtcZoneConcreteID       string
	MtcSectorConcreteID     string
	ChannelCode             string
	DataDirection           string
	Status                  string
	Extension               string
	ChannelShowMappingID    string
	ChannelSessionMappingID string
	ZoneConcreteName        string
	EndNo                   string
	RowName                 string
	ZoneConcreteCode        string
	FloorNo                 string
	CreateTime              string
	UpdateTime              string
	IsDeleted               string
}{
	ChannelMappingID:        "channel_mapping_id",
	ChannelRowConcreteID:    "channel_row_concrete_id",
	ChannelZoneConcreteID:   "channel_zone_concrete_id",
	MtcZoneConcreteID:       "mtc_zone_concrete_id",
	MtcSectorConcreteID:     "mtc_sector_concrete_id",
	ChannelCode:             "channel_code",
	DataDirection:           "data_direction",
	Status:                  "status",
	Extension:               "extension",
	ChannelShowMappingID:    "channel_show_mapping_id",
	ChannelSessionMappingID: "channel_session_mapping_id",
	ZoneConcreteName:        "zone_concrete_name",
	EndNo:                   "end_no",
	RowName:                 "row_name",
	ZoneConcreteCode:        "zone_concrete_code",
	FloorNo:                 "floor_no",
	CreateTime:              "create_time",
	UpdateTime:              "update_time",
	IsDeleted:               "is_deleted",
}

// BgChannelStandardZoneIDMappings [...]
type BgChannelStandardZoneIDMappings struct {
	ChannelMappingID             string `gorm:"primaryKey;column:channel_mapping_id" json:"-"`
	ChannelShowMappingID         string `gorm:"column:channel_show_mapping_id" json:"channelShowMappingId"`       // 演出映射Id
	ChannelSessionMappingID      string `gorm:"column:channel_session_mapping_id" json:"channelSessionMappingId"` // 场次映射Id
	ChannelZoneConcreteMappingID string `gorm:"column:channel_zone_concrete_mapping_id" json:"channelZoneConcreteMappingId"`
	ChannelCode                  string `gorm:"column:channel_code" json:"channelCode"` // 渠道编码
	MappingValue                 string `gorm:"column:mapping_value" json:"mappingValue"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelStandardZoneIDMappings) TableName() string {
	return "bg_channel_standard_zone_id_mappings"
}

// BgChannelStandardZoneIDMappingsColumns get sql column name.获取数据库列名
var BgChannelStandardZoneIDMappingsColumns = struct {
	ChannelMappingID             string
	ChannelShowMappingID         string
	ChannelSessionMappingID      string
	ChannelZoneConcreteMappingID string
	ChannelCode                  string
	MappingValue                 string
}{
	ChannelMappingID:             "channel_mapping_id",
	ChannelShowMappingID:         "channel_show_mapping_id",
	ChannelSessionMappingID:      "channel_session_mapping_id",
	ChannelZoneConcreteMappingID: "channel_zone_concrete_mapping_id",
	ChannelCode:                  "channel_code",
	MappingValue:                 "mapping_value",
}

// BgChannelState [...]
type BgChannelState struct {
	ID          int       `gorm:"primaryKey;column:id" json:"-"`
	ChannelType string    `gorm:"column:channel_type" json:"channelType"` // 渠道类型
	Channel     string    `gorm:"column:channel" json:"channel"`          // 渠道名称
	State       []uint8   `gorm:"column:state" json:"state"`              // 渠道状态：0：关闭 1：开启
	UpdateTime  time.Time `gorm:"column:updateTime" json:"updateTime"`
	CreateTime  time.Time `gorm:"column:createTime" json:"createTime"`
	IsDeleted   []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelState) TableName() string {
	return "bg_channel_state"
}

// BgChannelStateColumns get sql column name.获取数据库列名
var BgChannelStateColumns = struct {
	ID          string
	ChannelType string
	Channel     string
	State       string
	UpdateTime  string
	CreateTime  string
	IsDeleted   string
}{
	ID:          "id",
	ChannelType: "channel_type",
	Channel:     "channel",
	State:       "state",
	UpdateTime:  "updateTime",
	CreateTime:  "createTime",
	IsDeleted:   "isDeleted",
}

// BgChannelSupplierMappingHistory [...]
type BgChannelSupplierMappingHistory struct {
	ID            int       `gorm:"primaryKey;column:id" json:"-"`
	ResourceTable string    `gorm:"column:resourceTable" json:"resourceTable"` // 来源映射表
	MappingID     string    `gorm:"column:mappingId" json:"mappingId"`         // 映射表Id
	MappingType   string    `gorm:"column:mappingType" json:"mappingType"`     // 映射类型 三类 演出，场次，票面
	SupplierID    string    `gorm:"column:supplierId" json:"supplierId"`       // 供应商
	SystemID      string    `gorm:"column:systemId" json:"systemId"`           // 平台id
	Reason        string    `gorm:"column:reason" json:"reason"`               // 关系解除原因
	Channel       string    `gorm:"column:channel" json:"channel"`             // 渠道
	CreateTime    time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime    time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted     []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelSupplierMappingHistory) TableName() string {
	return "bg_channel_supplier_mapping_history"
}

// BgChannelSupplierMappingHistoryColumns get sql column name.获取数据库列名
var BgChannelSupplierMappingHistoryColumns = struct {
	ID            string
	ResourceTable string
	MappingID     string
	MappingType   string
	SupplierID    string
	SystemID      string
	Reason        string
	Channel       string
	CreateTime    string
	UpdateTime    string
	IsDeleted     string
}{
	ID:            "id",
	ResourceTable: "resourceTable",
	MappingID:     "mappingId",
	MappingType:   "mappingType",
	SupplierID:    "supplierId",
	SystemID:      "systemId",
	Reason:        "reason",
	Channel:       "channel",
	CreateTime:    "createTime",
	UpdateTime:    "updateTime",
	IsDeleted:     "isDeleted",
}

// BgChannelSupplierSessionIDMappings [...]
type BgChannelSupplierSessionIDMappings struct {
	SupplierSessionID string    `gorm:"primaryKey;column:supplier_session_id" json:"-"`
	MtcSessionID      string    `gorm:"column:mtc_session_id" json:"mtcSessionId"`
	CreateTime        time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime        time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted         []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelSupplierSessionIDMappings) TableName() string {
	return "bg_channel_supplier_session_id_mappings"
}

// BgChannelSupplierSessionIDMappingsColumns get sql column name.获取数据库列名
var BgChannelSupplierSessionIDMappingsColumns = struct {
	SupplierSessionID string
	MtcSessionID      string
	CreateTime        string
	UpdateTime        string
	IsDeleted         string
}{
	SupplierSessionID: "supplier_session_id",
	MtcSessionID:      "mtc_session_id",
	CreateTime:        "create_time",
	UpdateTime:        "update_time",
	IsDeleted:         "is_deleted",
}

// BgChannelSupplierShowIDMappings [...]
type BgChannelSupplierShowIDMappings struct {
	SupplierShowID string    `gorm:"primaryKey;column:supplier_show_id" json:"-"`
	MtcShowID      string    `gorm:"column:mtc_show_id" json:"mtcShowId"`
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted      []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelSupplierShowIDMappings) TableName() string {
	return "bg_channel_supplier_show_id_mappings"
}

// BgChannelSupplierShowIDMappingsColumns get sql column name.获取数据库列名
var BgChannelSupplierShowIDMappingsColumns = struct {
	SupplierShowID string
	MtcShowID      string
	CreateTime     string
	UpdateTime     string
	IsDeleted      string
}{
	SupplierShowID: "supplier_show_id",
	MtcShowID:      "mtc_show_id",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
	IsDeleted:      "is_deleted",
}

// BgChannelSupplierVenueFloorRegionTemp [...]
type BgChannelSupplierVenueFloorRegionTemp struct {
	ID                    int       `gorm:"primaryKey;column:id" json:"-"` // 映射Id
	ChannelRegionID       string    `gorm:"column:channel_region_id" json:"channelRegionId"`
	ChannelRowID          string    `gorm:"column:channel_row_id" json:"channelRowId"`
	ChannelCode           string    `gorm:"column:channel_code" json:"channelCode"` // 渠道编码
	ChannelVenueMappingID string    `gorm:"column:channel_venue_mapping_id" json:"channelVenueMappingId"`
	EndNo                 string    `gorm:"column:end_no" json:"endNo"` // 结束座位号
	ChannelRegionName     string    `gorm:"column:channel_region_name" json:"channelRegionName"`
	ChannelRowName        string    `gorm:"column:channel_row_name" json:"channelRowName"` // 排名称
	FloorNo               int       `gorm:"column:floor_no" json:"floorNo"`                // 场馆楼层模板Id
	CreateTime            time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime            time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted             []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelSupplierVenueFloorRegionTemp) TableName() string {
	return "bg_channel_supplier_venue_floor_region_temp"
}

// BgChannelSupplierVenueFloorRegionTempColumns get sql column name.获取数据库列名
var BgChannelSupplierVenueFloorRegionTempColumns = struct {
	ID                    string
	ChannelRegionID       string
	ChannelRowID          string
	ChannelCode           string
	ChannelVenueMappingID string
	EndNo                 string
	ChannelRegionName     string
	ChannelRowName        string
	FloorNo               string
	CreateTime            string
	UpdateTime            string
	IsDeleted             string
}{
	ID:                    "id",
	ChannelRegionID:       "channel_region_id",
	ChannelRowID:          "channel_row_id",
	ChannelCode:           "channel_code",
	ChannelVenueMappingID: "channel_venue_mapping_id",
	EndNo:                 "end_no",
	ChannelRegionName:     "channel_region_name",
	ChannelRowName:        "channel_row_name",
	FloorNo:               "floor_no",
	CreateTime:            "create_time",
	UpdateTime:            "update_time",
	IsDeleted:             "is_deleted",
}

// BgChannelSupportSpecialCategorys [...]
type BgChannelSupportSpecialCategorys struct {
	ID              string    `gorm:"primaryKey;column:id" json:"-"`
	Channel         string    `gorm:"column:channel" json:"channel"`
	SpecialCategory string    `gorm:"column:special_category" json:"specialCategory"`
	IsDeleted       []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
	UpdateTime      time.Time `gorm:"column:updateTime" json:"updateTime"`
	CreateTime      time.Time `gorm:"column:createTime" json:"createTime"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelSupportSpecialCategorys) TableName() string {
	return "bg_channel_support_special_categorys"
}

// BgChannelSupportSpecialCategorysColumns get sql column name.获取数据库列名
var BgChannelSupportSpecialCategorysColumns = struct {
	ID              string
	Channel         string
	SpecialCategory string
	IsDeleted       string
	UpdateTime      string
	CreateTime      string
}{
	ID:              "id",
	Channel:         "channel",
	SpecialCategory: "special_category",
	IsDeleted:       "isDeleted",
	UpdateTime:      "updateTime",
	CreateTime:      "createTime",
}

// BgChannelTestresultdata 推送消息测试结果记录表
type BgChannelTestresultdata struct {
	ResultDataID  string    `gorm:"primaryKey;column:result_data_id" json:"-"` // 主键
	ResourceID    string    `gorm:"column:resource_id" json:"resourceId"`      // 演出/场次/票面id
	Operation     string    `gorm:"column:operation" json:"operation"`         // 操作号
	Channel       string    `gorm:"column:channel" json:"channel"`             // 渠道名
	ResultMessage string    `gorm:"column:resultMessage" json:"resultMessage"` // 推送请求日志
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted     []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelTestresultdata) TableName() string {
	return "bg_channel_testresultdata"
}

// BgChannelTestresultdataColumns get sql column name.获取数据库列名
var BgChannelTestresultdataColumns = struct {
	ResultDataID  string
	ResourceID    string
	Operation     string
	Channel       string
	ResultMessage string
	CreateTime    string
	UpdateTime    string
	IsDeleted     string
}{
	ResultDataID:  "result_data_id",
	ResourceID:    "resource_id",
	Operation:     "operation",
	Channel:       "channel",
	ResultMessage: "resultMessage",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
	IsDeleted:     "is_deleted",
}

// BgChannelUncludeShowids [...]
type BgChannelUncludeShowids struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`
	Channel    string    `gorm:"column:channel" json:"channel"`
	ShowID     string    `gorm:"column:showId" json:"showId"`
	CreateTime time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted  []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelUncludeShowids) TableName() string {
	return "bg_channel_unclude_showids"
}

// BgChannelUncludeShowidsColumns get sql column name.获取数据库列名
var BgChannelUncludeShowidsColumns = struct {
	ID         string
	Channel    string
	ShowID     string
	CreateTime string
	UpdateTime string
	IsDeleted  string
}{
	ID:         "id",
	Channel:    "channel",
	ShowID:     "showId",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	IsDeleted:  "is_deleted",
}

// BgChannelURLMappings [...]
type BgChannelURLMappings struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`
	Channel   string `gorm:"column:channel" json:"channel"`
	SrcURL    string `gorm:"column:src_url" json:"srcUrl"`
	TargetURL string `gorm:"column:target_url" json:"targetUrl"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelURLMappings) TableName() string {
	return "bg_channel_url_mappings"
}

// BgChannelURLMappingsColumns get sql column name.获取数据库列名
var BgChannelURLMappingsColumns = struct {
	ID        string
	Channel   string
	SrcURL    string
	TargetURL string
}{
	ID:        "id",
	Channel:   "channel",
	SrcURL:    "src_url",
	TargetURL: "target_url",
}

// BgChannelUser [...]
type BgChannelUser struct {
	Oid        string  `gorm:"primaryKey;column:oid" json:"-"`
	Active     []uint8 `gorm:"column:active" json:"active"`
	ChannelOid string  `gorm:"column:channel_oid" json:"channelOid"`
	Login      string  `gorm:"column:login" json:"login"`
	Mobile     string  `gorm:"column:mobile" json:"mobile"`
	Password   string  `gorm:"column:password" json:"password"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelUser) TableName() string {
	return "bg_channel_user"
}

// BgChannelUserColumns get sql column name.获取数据库列名
var BgChannelUserColumns = struct {
	Oid        string
	Active     string
	ChannelOid string
	Login      string
	Mobile     string
	Password   string
}{
	Oid:        "oid",
	Active:     "active",
	ChannelOid: "channel_oid",
	Login:      "login",
	Mobile:     "mobile",
	Password:   "password",
}

// BgChannelVenueMapping 渠道场馆映射表
type BgChannelVenueMapping struct {
	ID                   uint64    `gorm:"primaryKey;column:id" json:"-"`
	LocalVenueID         string    `gorm:"column:localVenueId" json:"localVenueId"`                 // 系统场馆Id
	LocalVenueName       string    `gorm:"column:localVenueName" json:"localVenueName"`             // 系统场馆名称
	ChannelVenueID       string    `gorm:"column:channelVenueId" json:"channelVenueId"`             // 渠道场馆id
	ChannelVenueName     string    `gorm:"column:channelVenueName" json:"channelVenueName"`         // 渠道场馆名称
	ChannelName          string    `gorm:"column:channelName" json:"channelName"`                   // 渠道名称
	ChannelCityMappingID uint64    `gorm:"column:channelCityMappingId" json:"channelCityMappingId"` // 渠道城市映射id
	Enable               bool      `gorm:"column:enable" json:"enable"`                             // 映射配置的状态 1有效，0无效
	CreateTime           time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime           time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted            []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgChannelVenueMapping) TableName() string {
	return "bg_channel_venue_mapping"
}

// BgChannelVenueMappingColumns get sql column name.获取数据库列名
var BgChannelVenueMappingColumns = struct {
	ID                   string
	LocalVenueID         string
	LocalVenueName       string
	ChannelVenueID       string
	ChannelVenueName     string
	ChannelName          string
	ChannelCityMappingID string
	Enable               string
	CreateTime           string
	UpdateTime           string
	IsDeleted            string
}{
	ID:                   "id",
	LocalVenueID:         "localVenueId",
	LocalVenueName:       "localVenueName",
	ChannelVenueID:       "channelVenueId",
	ChannelVenueName:     "channelVenueName",
	ChannelName:          "channelName",
	ChannelCityMappingID: "channelCityMappingId",
	Enable:               "enable",
	CreateTime:           "createTime",
	UpdateTime:           "updateTime",
	IsDeleted:            "isDeleted",
}

// BgJdOrderTicket 京东订单票价id关系表
type BgJdOrderTicket struct {
	JdOrderID  string    `gorm:"primaryKey;column:jd_order_id" json:"-"` // 京东订单id
	JdTicketID string    `gorm:"column:jd_ticket_id" json:"jdTicketId"`  // 京东票价id
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted  []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgJdOrderTicket) TableName() string {
	return "bg_jd_order_ticket"
}

// BgJdOrderTicketColumns get sql column name.获取数据库列名
var BgJdOrderTicketColumns = struct {
	JdOrderID  string
	JdTicketID string
	CreateTime string
	UpdateTime string
	IsDeleted  string
}{
	JdOrderID:  "jd_order_id",
	JdTicketID: "jd_ticket_id",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	IsDeleted:  "is_deleted",
}

// BgJdSeatplanTicketMapping 京东票价id与票面关系表
type BgJdSeatplanTicketMapping struct {
	JdTicketID string    `gorm:"primaryKey;column:jd_ticket_id" json:"-"` // 京东票价id
	SeatplanID string    `gorm:"column:seatplan_id" json:"seatplanId"`    // 票面id
	IsHidden   []uint8   `gorm:"column:is_hidden" json:"isHidden"`        // 是否隐藏
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted  []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgJdSeatplanTicketMapping) TableName() string {
	return "bg_jd_seatplan_ticket_mapping"
}

// BgJdSeatplanTicketMappingColumns get sql column name.获取数据库列名
var BgJdSeatplanTicketMappingColumns = struct {
	JdTicketID string
	SeatplanID string
	IsHidden   string
	CreateTime string
	UpdateTime string
	IsDeleted  string
}{
	JdTicketID: "jd_ticket_id",
	SeatplanID: "seatplan_id",
	IsHidden:   "is_hidden",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	IsDeleted:  "is_deleted",
}

// BgOrderCenterOrderIDMapping [...]
type BgOrderCenterOrderIDMapping struct {
	ID             uint32    `gorm:"primaryKey;column:id" json:"-"`
	MtcOrderNumber string    `gorm:"column:mtc_order_number" json:"mtcOrderNumber"` // 订单号
	Channel        string    `gorm:"column:channel" json:"channel"`                 // 渠道
	ChannelOrderID string    `gorm:"column:channel_order_id" json:"channelOrderId"` // 渠道订单号
	Sort           int       `gorm:"column:sort" json:"sort"`                       // 排序
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted      []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgOrderCenterOrderIDMapping) TableName() string {
	return "bg_order_center_order_id_mapping"
}

// BgOrderCenterOrderIDMappingColumns get sql column name.获取数据库列名
var BgOrderCenterOrderIDMappingColumns = struct {
	ID             string
	MtcOrderNumber string
	Channel        string
	ChannelOrderID string
	Sort           string
	CreateTime     string
	UpdateTime     string
	IsDeleted      string
}{
	ID:             "id",
	MtcOrderNumber: "mtc_order_number",
	Channel:        "channel",
	ChannelOrderID: "channel_order_id",
	Sort:           "sort",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
	IsDeleted:      "is_deleted",
}

// BgOrderWithUnmatchedStatus 渠道订单状态异常记录
type BgOrderWithUnmatchedStatus struct {
	OrderID            string    `gorm:"primaryKey;column:order_id" json:"-"`                   // 订单id
	OrderStatus        string    `gorm:"column:order_status" json:"orderStatus"`                // 订单状态
	ChannelOrderID     string    `gorm:"column:channel_order_id" json:"channelOrderId"`         // 渠道订单id
	ChannelOrderStatus string    `gorm:"column:channel_order_status" json:"channelOrderStatus"` // 渠道订单状态
	Channel            string    `gorm:"column:channel" json:"channel"`                         // 渠道编号
	CreateTime         time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime         time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted          []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgOrderWithUnmatchedStatus) TableName() string {
	return "bg_order_with_unmatched_status"
}

// BgOrderWithUnmatchedStatusColumns get sql column name.获取数据库列名
var BgOrderWithUnmatchedStatusColumns = struct {
	OrderID            string
	OrderStatus        string
	ChannelOrderID     string
	ChannelOrderStatus string
	Channel            string
	CreateTime         string
	UpdateTime         string
	IsDeleted          string
}{
	OrderID:            "order_id",
	OrderStatus:        "order_status",
	ChannelOrderID:     "channel_order_id",
	ChannelOrderStatus: "channel_order_status",
	Channel:            "channel",
	CreateTime:         "create_time",
	UpdateTime:         "update_time",
	IsDeleted:          "is_deleted",
}

// BgResourceIDMappings 渠道资源ID映射关系（资源为演出、排期、票面等）
type BgResourceIDMappings struct {
	MappingID               string    `gorm:"primaryKey;column:mapping_id" json:"-"`
	ResourceID              string    `gorm:"column:resource_id" json:"resourceId"`                             // 摩天轮侧ID
	ChannelResourceID       string    `gorm:"column:channel_resource_id" json:"channelResourceId"`              // 渠道侧ID
	ChannelParentResourceID string    `gorm:"column:channel_parent_resource_id" json:"channelParentResourceId"` // 渠道侧父ID
	OldChannelStatus        bool      `gorm:"column:old_channel_status" json:"oldChannelStatus"`                // 渠道侧上次状态
	ChannelStatus           bool      `gorm:"column:channel_status" json:"channelStatus"`                       // 渠道侧状态
	Status                  bool      `gorm:"column:status" json:"status"`                                      // 演出状态 1 预售中 2 售票中; 排期状态　1 待售 2 销售中 3 售罄; 票面状态 可售 售完
	ResourceType            int8      `gorm:"column:resource_type" json:"resourceType"`                         // 摩天轮资源类型，1 演出 2 排期 3 票面 4 票
	Channel                 string    `gorm:"column:channel" json:"channel"`                                    // 渠道编号
	ChannelResourceType     int8      `gorm:"column:channel_resource_type" json:"channelResourceType"`          // 渠道资源类型 1 演出 2 排期 3 票面 4 票
	CreateTime              time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime              time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted               bool      `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgResourceIDMappings) TableName() string {
	return "bg_resource_id_mappings"
}

// BgResourceIDMappingsColumns get sql column name.获取数据库列名
var BgResourceIDMappingsColumns = struct {
	MappingID               string
	ResourceID              string
	ChannelResourceID       string
	ChannelParentResourceID string
	OldChannelStatus        string
	ChannelStatus           string
	Status                  string
	ResourceType            string
	Channel                 string
	ChannelResourceType     string
	CreateTime              string
	UpdateTime              string
	IsDeleted               string
}{
	MappingID:               "mapping_id",
	ResourceID:              "resource_id",
	ChannelResourceID:       "channel_resource_id",
	ChannelParentResourceID: "channel_parent_resource_id",
	OldChannelStatus:        "old_channel_status",
	ChannelStatus:           "channel_status",
	Status:                  "status",
	ResourceType:            "resource_type",
	Channel:                 "channel",
	ChannelResourceType:     "channel_resource_type",
	CreateTime:              "create_time",
	UpdateTime:              "update_time",
	IsDeleted:               "is_deleted",
}

// BgSeatPlanAudit 渠道票面同步待审核
type BgSeatPlanAudit struct {
	SeatPlanAuditID    string    `gorm:"primaryKey;column:seat_plan_audit_id" json:"-"`
	ShowAuditID        string    `gorm:"column:show_audit_id" json:"showAuditId"`                // 演出审核记录ID
	ShowSessionAuditID string    `gorm:"column:show_session_audit_id" json:"showSessionAuditId"` // 排期审核记录ID
	AuditStatus        []uint8   `gorm:"column:audit_status" json:"auditStatus"`                 // 审核状态 1 待审核 2 已审核 3 审核不通过
	Channel            string    `gorm:"column:channel" json:"channel"`                          // 渠道编号
	AuditComment       string    `gorm:"column:audit_comment" json:"auditComment"`               // 审核描述
	OriginalPrice      float64   `gorm:"column:original_price" json:"originalPrice"`             // 票价
	Comments           string    `gorm:"column:comments" json:"comments"`                        // 说明
	SeatPlanUnit       string    `gorm:"column:seat_plan_unit" json:"seatPlanUnit"`              // 普通票或套票
	UnitQty            int       `gorm:"column:unit_qty" json:"unitQty"`                         // 套票张数
	CreateTime         time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime         time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted          []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgSeatPlanAudit) TableName() string {
	return "bg_seat_plan_audit"
}

// BgSeatPlanAuditColumns get sql column name.获取数据库列名
var BgSeatPlanAuditColumns = struct {
	SeatPlanAuditID    string
	ShowAuditID        string
	ShowSessionAuditID string
	AuditStatus        string
	Channel            string
	AuditComment       string
	OriginalPrice      string
	Comments           string
	SeatPlanUnit       string
	UnitQty            string
	CreateTime         string
	UpdateTime         string
	IsDeleted          string
}{
	SeatPlanAuditID:    "seat_plan_audit_id",
	ShowAuditID:        "show_audit_id",
	ShowSessionAuditID: "show_session_audit_id",
	AuditStatus:        "audit_status",
	Channel:            "channel",
	AuditComment:       "audit_comment",
	OriginalPrice:      "original_price",
	Comments:           "comments",
	SeatPlanUnit:       "seat_plan_unit",
	UnitQty:            "unit_qty",
	CreateTime:         "create_time",
	UpdateTime:         "update_time",
	IsDeleted:          "is_deleted",
}

// BgShowAudit 渠道演出同步待审核
type BgShowAudit struct {
	ShowAuditID    string    `gorm:"primaryKey;column:show_audit_id" json:"-"` // 演出审核记录ID
	ShowName       string    `gorm:"column:show_name" json:"showName"`         // 演出名称
	Advertise      string    `gorm:"column:advertise" json:"advertise"`
	ShowType       string    `gorm:"column:show_type" json:"showType"`
	AuditStatus    []uint8   `gorm:"column:audit_status" json:"auditStatus"`       // 演出审核状态 1 待审核 2 已审核 3 审核不通过
	AuditComment   string    `gorm:"column:audit_comment" json:"auditComment"`     // 审核描述
	ShowDesc       string    `gorm:"column:show_desc" json:"showDesc"`             // 演出介绍
	Content        string    `gorm:"column:content" json:"content"`                // 演出内容
	ShowStatus     bool      `gorm:"column:show_status" json:"showStatus"`         // 演出状态 1 预售中 2 售票中
	PosterURL      string    `gorm:"column:poster_url" json:"posterUrl"`           // 演出海报
	SeatPlanURL    string    `gorm:"column:seat_plan_url" json:"seatPlanUrl"`      // 座位图
	VenueID        string    `gorm:"column:venue_id" json:"venueId"`               // 场馆ID
	VenueName      string    `gorm:"column:venue_name" json:"venueName"`           // 场馆名称
	CityID         int       `gorm:"column:city_id" json:"cityId"`                 // 城市ID
	CityName       string    `gorm:"column:city_name" json:"cityName"`             // 城市名称
	Channel        string    `gorm:"column:channel" json:"channel"`                // 渠道编号
	SupportExpress []uint8   `gorm:"column:support_express" json:"supportExpress"` // 支持快递
	SupportVenue   []uint8   `gorm:"column:support_venue" json:"supportVenue"`     // 支持现场取票
	SupportOnsite  []uint8   `gorm:"column:support_onsite" json:"supportOnsite"`   // 支持上门取票
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted      []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgShowAudit) TableName() string {
	return "bg_show_audit"
}

// BgShowAuditColumns get sql column name.获取数据库列名
var BgShowAuditColumns = struct {
	ShowAuditID    string
	ShowName       string
	Advertise      string
	ShowType       string
	AuditStatus    string
	AuditComment   string
	ShowDesc       string
	Content        string
	ShowStatus     string
	PosterURL      string
	SeatPlanURL    string
	VenueID        string
	VenueName      string
	CityID         string
	CityName       string
	Channel        string
	SupportExpress string
	SupportVenue   string
	SupportOnsite  string
	CreateTime     string
	UpdateTime     string
	IsDeleted      string
}{
	ShowAuditID:    "show_audit_id",
	ShowName:       "show_name",
	Advertise:      "advertise",
	ShowType:       "show_type",
	AuditStatus:    "audit_status",
	AuditComment:   "audit_comment",
	ShowDesc:       "show_desc",
	Content:        "content",
	ShowStatus:     "show_status",
	PosterURL:      "poster_url",
	SeatPlanURL:    "seat_plan_url",
	VenueID:        "venue_id",
	VenueName:      "venue_name",
	CityID:         "city_id",
	CityName:       "city_name",
	Channel:        "channel",
	SupportExpress: "support_express",
	SupportVenue:   "support_venue",
	SupportOnsite:  "support_onsite",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
	IsDeleted:      "is_deleted",
}

// BgShowSessionAudit 渠道排期同步待审核
type BgShowSessionAudit struct {
	ShowSessionAuditID string    `gorm:"primaryKey;column:show_session_audit_id" json:"-"` // 排期审核记录ID
	SessionName        string    `gorm:"column:session_name" json:"sessionName"`           // 排期名称
	ShowAuditID        string    `gorm:"column:show_audit_id" json:"showAuditId"`          // 演出审核记录ID
	ShowTime           time.Time `gorm:"column:show_time" json:"showTime"`                 // 演出时间
	AuditStatus        []uint8   `gorm:"column:audit_status" json:"auditStatus"`           // 审核状态 1 待审核 2 已审核 3 审核不通过
	AuditComment       string    `gorm:"column:audit_comment" json:"auditComment"`         // 审核描述
	Channel            string    `gorm:"column:channel" json:"channel"`                    // 渠道编号
	CreateTime         time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime         time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted          []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgShowSessionAudit) TableName() string {
	return "bg_show_session_audit"
}

// BgShowSessionAuditColumns get sql column name.获取数据库列名
var BgShowSessionAuditColumns = struct {
	ShowSessionAuditID string
	SessionName        string
	ShowAuditID        string
	ShowTime           string
	AuditStatus        string
	AuditComment       string
	Channel            string
	CreateTime         string
	UpdateTime         string
	IsDeleted          string
}{
	ShowSessionAuditID: "show_session_audit_id",
	SessionName:        "session_name",
	ShowAuditID:        "show_audit_id",
	ShowTime:           "show_time",
	AuditStatus:        "audit_status",
	AuditComment:       "audit_comment",
	Channel:            "channel",
	CreateTime:         "create_time",
	UpdateTime:         "update_time",
	IsDeleted:          "is_deleted",
}

// BgSupplierSeatPlan 供应商票面数据映射关系
type BgSupplierSeatPlan struct {
	MappingID             string    `gorm:"primaryKey;column:mapping_id" json:"-"`                  // 主键
	SupplierSeatPlanID    string    `gorm:"column:supplier_seat_plan_id" json:"supplierSeatPlanId"` // 供应商票面id
	Comments              string    `gorm:"column:comments" json:"comments"`                        // 演出时间
	Enabled               []uint8   `gorm:"column:enabled" json:"enabled"`                          // 演出状态
	Price                 float64   `gorm:"column:price" json:"price"`
	Limitation            int       `gorm:"column:limitation" json:"limitation"` // 限购值
	SeatPlanUnit          string    `gorm:"column:seat_plan_unit" json:"seatPlanUnit"`
	SupportEticket        []uint8   `gorm:"column:support_eticket" json:"supportEticket"` // 是否支持电子票
	UnitQty               int       `gorm:"column:unit_qty" json:"unitQty"`
	SupplierShowID        string    `gorm:"column:supplier_show_id" json:"supplierShowId"`                // 供应商演出id
	SupplierShowSessionID string    `gorm:"column:supplier_show_session_id" json:"supplierShowSessionId"` // 供应商排期id
	ShowSessionID         string    `gorm:"column:show_session_id" json:"showSessionId"`                  // 排期id
	ShowID                string    `gorm:"column:show_id" json:"showId"`                                 // 演出id
	SeatPlanID            string    `gorm:"column:seat_plan_id" json:"seatPlanId"`                        // 票面id
	Version               int64     `gorm:"column:version" json:"version"`                                // 票面版本
	Channel               string    `gorm:"column:channel" json:"channel"`                                // 渠道
	NotMatchReason        string    `gorm:"column:not_match_reason" json:"notMatchReason"`                // 映射关系未匹配上的原因
	CreateTime            time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime            time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted             []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgSupplierSeatPlan) TableName() string {
	return "bg_supplier_seat_plan"
}

// BgSupplierSeatPlanColumns get sql column name.获取数据库列名
var BgSupplierSeatPlanColumns = struct {
	MappingID             string
	SupplierSeatPlanID    string
	Comments              string
	Enabled               string
	Price                 string
	Limitation            string
	SeatPlanUnit          string
	SupportEticket        string
	UnitQty               string
	SupplierShowID        string
	SupplierShowSessionID string
	ShowSessionID         string
	ShowID                string
	SeatPlanID            string
	Version               string
	Channel               string
	NotMatchReason        string
	CreateTime            string
	UpdateTime            string
	IsDeleted             string
}{
	MappingID:             "mapping_id",
	SupplierSeatPlanID:    "supplier_seat_plan_id",
	Comments:              "comments",
	Enabled:               "enabled",
	Price:                 "price",
	Limitation:            "limitation",
	SeatPlanUnit:          "seat_plan_unit",
	SupportEticket:        "support_eticket",
	UnitQty:               "unit_qty",
	SupplierShowID:        "supplier_show_id",
	SupplierShowSessionID: "supplier_show_session_id",
	ShowSessionID:         "show_session_id",
	ShowID:                "show_id",
	SeatPlanID:            "seat_plan_id",
	Version:               "version",
	Channel:               "channel",
	NotMatchReason:        "not_match_reason",
	CreateTime:            "create_time",
	UpdateTime:            "update_time",
	IsDeleted:             "is_deleted",
}

// BgSupplierShow 供应商演出数据映射关系
type BgSupplierShow struct {
	MappingID       string    `gorm:"primaryKey;column:mapping_id" json:"-"`           // 主键
	SupplierShowID  string    `gorm:"column:supplier_show_id" json:"supplierShowId"`   // 供应商演出id
	ShowName        string    `gorm:"column:show_name" json:"showName"`                // 演出名
	ShowType        string    `gorm:"column:show_type" json:"showType"`                // 演出类型
	ShowStatus      string    `gorm:"column:show_status" json:"showStatus"`            // 演出状态
	SupplierVenueID string    `gorm:"column:supplier_venue_id" json:"supplierVenueId"` // 供应商场馆id
	LocalVenueID    string    `gorm:"column:local_venue_id" json:"localVenueId"`
	VenueName       string    `gorm:"column:venue_name" json:"venueName"`             // 场馆名
	SupportDelivery string    `gorm:"column:support_delivery" json:"supportDelivery"` // 支持的配送方式，按2进程数表示，0不支持，1支持，第1位为快递，第2位为上门，第3位为现场
	CityID          string    `gorm:"column:city_id" json:"cityId"`                   // 城市id
	Limitation      int       `gorm:"column:limitation" json:"limitation"`            // 限购值
	CardSuggestType int8      `gorm:"column:card_suggest_type" json:"cardSuggestType"`
	ShowID          string    `gorm:"column:show_id" json:"showId"`  // 演出id
	Channel         string    `gorm:"column:channel" json:"channel"` // 渠道
	CreateTime      time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted       []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgSupplierShow) TableName() string {
	return "bg_supplier_show"
}

// BgSupplierShowColumns get sql column name.获取数据库列名
var BgSupplierShowColumns = struct {
	MappingID       string
	SupplierShowID  string
	ShowName        string
	ShowType        string
	ShowStatus      string
	SupplierVenueID string
	LocalVenueID    string
	VenueName       string
	SupportDelivery string
	CityID          string
	Limitation      string
	CardSuggestType string
	ShowID          string
	Channel         string
	CreateTime      string
	UpdateTime      string
	IsDeleted       string
}{
	MappingID:       "mapping_id",
	SupplierShowID:  "supplier_show_id",
	ShowName:        "show_name",
	ShowType:        "show_type",
	ShowStatus:      "show_status",
	SupplierVenueID: "supplier_venue_id",
	LocalVenueID:    "local_venue_id",
	VenueName:       "venue_name",
	SupportDelivery: "support_delivery",
	CityID:          "city_id",
	Limitation:      "limitation",
	CardSuggestType: "card_suggest_type",
	ShowID:          "show_id",
	Channel:         "channel",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
	IsDeleted:       "is_deleted",
}

// BgSupplierShow09251501 [...]
type BgSupplierShow09251501 struct {
	MappingID       string    `gorm:"column:mapping_id" json:"mappingId"`              // 主键
	SupplierShowID  string    `gorm:"column:supplier_show_id" json:"supplierShowId"`   // 供应商演出id
	ShowName        string    `gorm:"column:show_name" json:"showName"`                // 演出名
	ShowType        string    `gorm:"column:show_type" json:"showType"`                // 演出类型
	ShowStatus      string    `gorm:"column:show_status" json:"showStatus"`            // 演出状态
	SupplierVenueID string    `gorm:"column:supplier_venue_id" json:"supplierVenueId"` // 供应商场馆id
	LocalVenueID    string    `gorm:"column:local_venue_id" json:"localVenueId"`
	VenueName       string    `gorm:"column:venue_name" json:"venueName"`              // 场馆名
	SupportDelivery string    `gorm:"column:support_delivery" json:"supportDelivery"`  // 支持的配送方式，按2进程数表示，0不支持，1支持，第1位为快递，第2位为上门，第3位为现场
	CityID          string    `gorm:"column:city_id" json:"cityId"`                    // 城市id
	Limitation      int       `gorm:"column:limitation" json:"limitation"`             // 限购值
	ShowID          string    `gorm:"column:show_id" json:"showId"`                    // 演出id
	Channel         string    `gorm:"column:channel" json:"channel"`                   // 渠道
	CardSuggestType int8      `gorm:"column:card_suggest_type" json:"cardSuggestType"` // 演出身份证类型, 0:无需身份证，1:一票一证，2：一单一证
	CreateTime      time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted       []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgSupplierShow09251501) TableName() string {
	return "bg_supplier_show_09251501"
}

// BgSupplierShow09251501Columns get sql column name.获取数据库列名
var BgSupplierShow09251501Columns = struct {
	MappingID       string
	SupplierShowID  string
	ShowName        string
	ShowType        string
	ShowStatus      string
	SupplierVenueID string
	LocalVenueID    string
	VenueName       string
	SupportDelivery string
	CityID          string
	Limitation      string
	ShowID          string
	Channel         string
	CardSuggestType string
	CreateTime      string
	UpdateTime      string
	IsDeleted       string
}{
	MappingID:       "mapping_id",
	SupplierShowID:  "supplier_show_id",
	ShowName:        "show_name",
	ShowType:        "show_type",
	ShowStatus:      "show_status",
	SupplierVenueID: "supplier_venue_id",
	LocalVenueID:    "local_venue_id",
	VenueName:       "venue_name",
	SupportDelivery: "support_delivery",
	CityID:          "city_id",
	Limitation:      "limitation",
	ShowID:          "show_id",
	Channel:         "channel",
	CardSuggestType: "card_suggest_type",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
	IsDeleted:       "is_deleted",
}

// BgSupplierShowCentreMapping [...]
type BgSupplierShowCentreMapping struct {
	MappingID                int       `gorm:"primaryKey;column:mapping_id" json:"-"`
	SupplierResourceID       string    `gorm:"column:supplier_resource_id" json:"supplierResourceId"`              // 供应商资源Id
	SupplierResourceParentID string    `gorm:"column:supplier_resource_parent_id" json:"supplierResourceParentId"` // 供应商资源父Id
	ResourceType             string    `gorm:"column:resource_type" json:"resourceType"`                           // 资源类型1：show;2:session;3:seatPlan;4:ticket
	MtcResourceID            string    `gorm:"column:mtc_resource_id" json:"mtcResourceId"`                        // 演出中心资源ID
	Enable                   []uint8   `gorm:"column:enable" json:"enable"`                                        // 是否开启
	Supplier                 string    `gorm:"column:supplier" json:"supplier"`                                    // 供应商
	CreateTime               time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime               time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted                []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgSupplierShowCentreMapping) TableName() string {
	return "bg_supplier_show_centre_mapping"
}

// BgSupplierShowCentreMappingColumns get sql column name.获取数据库列名
var BgSupplierShowCentreMappingColumns = struct {
	MappingID                string
	SupplierResourceID       string
	SupplierResourceParentID string
	ResourceType             string
	MtcResourceID            string
	Enable                   string
	Supplier                 string
	CreateTime               string
	UpdateTime               string
	IsDeleted                string
}{
	MappingID:                "mapping_id",
	SupplierResourceID:       "supplier_resource_id",
	SupplierResourceParentID: "supplier_resource_parent_id",
	ResourceType:             "resource_type",
	MtcResourceID:            "mtc_resource_id",
	Enable:                   "enable",
	Supplier:                 "supplier",
	CreateTime:               "create_time",
	UpdateTime:               "update_time",
	IsDeleted:                "is_deleted",
}

// BgSupplierShowSession 供应商排期数据映射关系
type BgSupplierShowSession struct {
	MappingID             string    `gorm:"primaryKey;column:mapping_id" json:"-"`                        // 主键
	SupplierShowSessionID string    `gorm:"column:supplier_show_session_id" json:"supplierShowSessionId"` // 供应商排期id
	ShowTime              time.Time `gorm:"column:show_time" json:"showTime"`                             // 演出时间
	ShowStatus            string    `gorm:"column:show_status" json:"showStatus"`                         // 演出状态
	Limitation            int       `gorm:"column:limitation" json:"limitation"`                          // 限购值
	SupportDelivery       string    `gorm:"column:support_delivery" json:"supportDelivery"`               // 支持的配送方式，按2进程数表示，0不支持，1支持，第1位为快递，第2位为上门，第3位为现场
	IsSponsor             []uint8   `gorm:"column:is_sponsor" json:"isSponsor"`                           // 是否为主办 0 否 1 是
	SupplierShowID        string    `gorm:"column:supplier_show_id" json:"supplierShowId"`                // 供应商演出id
	ShowSessionID         string    `gorm:"column:show_session_id" json:"showSessionId"`                  // 排期id
	ShowID                string    `gorm:"column:show_id" json:"showId"`                                 // 演出id
	Channel               string    `gorm:"column:channel" json:"channel"`                                // 渠道
	CardSuggestType       int8      `gorm:"column:card_suggest_type" json:"cardSuggestType"`              // 演出身份证类型, 0:无需身份证，1:一票一证，2：一单一证
	CreateTime            time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime            time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted             []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgSupplierShowSession) TableName() string {
	return "bg_supplier_show_session"
}

// BgSupplierShowSessionColumns get sql column name.获取数据库列名
var BgSupplierShowSessionColumns = struct {
	MappingID             string
	SupplierShowSessionID string
	ShowTime              string
	ShowStatus            string
	Limitation            string
	SupportDelivery       string
	IsSponsor             string
	SupplierShowID        string
	ShowSessionID         string
	ShowID                string
	Channel               string
	CardSuggestType       string
	CreateTime            string
	UpdateTime            string
	IsDeleted             string
}{
	MappingID:             "mapping_id",
	SupplierShowSessionID: "supplier_show_session_id",
	ShowTime:              "show_time",
	ShowStatus:            "show_status",
	Limitation:            "limitation",
	SupportDelivery:       "support_delivery",
	IsSponsor:             "is_sponsor",
	SupplierShowID:        "supplier_show_id",
	ShowSessionID:         "show_session_id",
	ShowID:                "show_id",
	Channel:               "channel",
	CardSuggestType:       "card_suggest_type",
	CreateTime:            "create_time",
	UpdateTime:            "update_time",
	IsDeleted:             "is_deleted",
}

// BgSupplierShowSession09251501 [...]
type BgSupplierShowSession09251501 struct {
	MappingID             string    `gorm:"column:mapping_id" json:"mappingId"`                           // 主键
	SupplierShowSessionID string    `gorm:"column:supplier_show_session_id" json:"supplierShowSessionId"` // 供应商排期id
	ShowTime              time.Time `gorm:"column:show_time" json:"showTime"`                             // 演出时间
	ShowStatus            string    `gorm:"column:show_status" json:"showStatus"`                         // 演出状态
	Limitation            int       `gorm:"column:limitation" json:"limitation"`                          // 限购值
	SupportDelivery       string    `gorm:"column:support_delivery" json:"supportDelivery"`               // 支持的配送方式，按2进程数表示，0不支持，1支持，第1位为快递，第2位为上门，第3位为现场
	IsSponsor             []uint8   `gorm:"column:is_sponsor" json:"isSponsor"`                           // 是否为主办 0 否 1 是
	SupplierShowID        string    `gorm:"column:supplier_show_id" json:"supplierShowId"`                // 供应商演出id
	ShowSessionID         string    `gorm:"column:show_session_id" json:"showSessionId"`                  // 排期id
	ShowID                string    `gorm:"column:show_id" json:"showId"`                                 // 演出id
	Channel               string    `gorm:"column:channel" json:"channel"`                                // 渠道
	CardSuggestType       int8      `gorm:"column:card_suggest_type" json:"cardSuggestType"`              // 演出身份证类型, 0:无需身份证，1:一票一证，2：一单一证
	CreateTime            time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime            time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted             []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgSupplierShowSession09251501) TableName() string {
	return "bg_supplier_show_session_09251501"
}

// BgSupplierShowSession09251501Columns get sql column name.获取数据库列名
var BgSupplierShowSession09251501Columns = struct {
	MappingID             string
	SupplierShowSessionID string
	ShowTime              string
	ShowStatus            string
	Limitation            string
	SupportDelivery       string
	IsSponsor             string
	SupplierShowID        string
	ShowSessionID         string
	ShowID                string
	Channel               string
	CardSuggestType       string
	CreateTime            string
	UpdateTime            string
	IsDeleted             string
}{
	MappingID:             "mapping_id",
	SupplierShowSessionID: "supplier_show_session_id",
	ShowTime:              "show_time",
	ShowStatus:            "show_status",
	Limitation:            "limitation",
	SupportDelivery:       "support_delivery",
	IsSponsor:             "is_sponsor",
	SupplierShowID:        "supplier_show_id",
	ShowSessionID:         "show_session_id",
	ShowID:                "show_id",
	Channel:               "channel",
	CardSuggestType:       "card_suggest_type",
	CreateTime:            "create_time",
	UpdateTime:            "update_time",
	IsDeleted:             "is_deleted",
}

// BgSupplierSuccessMessage 渠道消息成功处理表
type BgSupplierSuccessMessage struct {
	MessageID  string    `gorm:"column:message_id" json:"messageId"` // 消息Id
	ID         string    `gorm:"primaryKey;column:id" json:"-"`      // 记录id
	Channel    string    `gorm:"column:channel" json:"channel"`      // 渠道
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted  []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgSupplierSuccessMessage) TableName() string {
	return "bg_supplier_success_message"
}

// BgSupplierSuccessMessageColumns get sql column name.获取数据库列名
var BgSupplierSuccessMessageColumns = struct {
	MessageID  string
	ID         string
	Channel    string
	CreateTime string
	UpdateTime string
	IsDeleted  string
}{
	MessageID:  "message_id",
	ID:         "id",
	Channel:    "channel",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	IsDeleted:  "is_deleted",
}

// BgSupplierTicket 供应商票数据映射关系
type BgSupplierTicket struct {
	MappingID             string    `gorm:"primaryKey;column:mapping_id" json:"-"`             // 主键
	SupplierTicketID      string    `gorm:"column:supplier_ticket_id" json:"supplierTicketId"` // 供应商票id
	Price                 float64   `gorm:"column:price" json:"price"`                         // 票价
	Stock                 int       `gorm:"column:stock" json:"stock"`
	TicketStatus          string    `gorm:"column:ticket_status" json:"ticketStatus"`                     // 票状态
	SupplierShowID        string    `gorm:"column:supplier_show_id" json:"supplierShowId"`                // 供应商演出id
	SupplierShowSessionID string    `gorm:"column:supplier_show_session_id" json:"supplierShowSessionId"` // 供应商排期id
	SupplierSeatPlanID    string    `gorm:"column:supplier_seat_plan_id" json:"supplierSeatPlanId"`       // 供应商票面id
	ShowSessionID         string    `gorm:"column:show_session_id" json:"showSessionId"`                  // 排期id
	ShowID                string    `gorm:"column:show_id" json:"showId"`                                 // 演出id
	SeatPlanID            string    `gorm:"column:seat_plan_id" json:"seatPlanId"`                        // 票面id
	TicketID              string    `gorm:"column:ticket_id" json:"ticketId"`                             // 票id
	SellerID              string    `gorm:"column:seller_id" json:"sellerId"`                             // 卖家id
	SynFlag               []uint8   `gorm:"column:syn_flag" json:"synFlag"`                               // 数据是否已同步到ticket:默认0未同步
	Channel               string    `gorm:"column:channel" json:"channel"`                                // 渠道
	CreateTime            time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime            time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted             []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgSupplierTicket) TableName() string {
	return "bg_supplier_ticket"
}

// BgSupplierTicketColumns get sql column name.获取数据库列名
var BgSupplierTicketColumns = struct {
	MappingID             string
	SupplierTicketID      string
	Price                 string
	Stock                 string
	TicketStatus          string
	SupplierShowID        string
	SupplierShowSessionID string
	SupplierSeatPlanID    string
	ShowSessionID         string
	ShowID                string
	SeatPlanID            string
	TicketID              string
	SellerID              string
	SynFlag               string
	Channel               string
	CreateTime            string
	UpdateTime            string
	IsDeleted             string
}{
	MappingID:             "mapping_id",
	SupplierTicketID:      "supplier_ticket_id",
	Price:                 "price",
	Stock:                 "stock",
	TicketStatus:          "ticket_status",
	SupplierShowID:        "supplier_show_id",
	SupplierShowSessionID: "supplier_show_session_id",
	SupplierSeatPlanID:    "supplier_seat_plan_id",
	ShowSessionID:         "show_session_id",
	ShowID:                "show_id",
	SeatPlanID:            "seat_plan_id",
	TicketID:              "ticket_id",
	SellerID:              "seller_id",
	SynFlag:               "syn_flag",
	Channel:               "channel",
	CreateTime:            "create_time",
	UpdateTime:            "update_time",
	IsDeleted:             "is_deleted",
}

// BgSyncExclution 无需数据同步条件配置表
type BgSyncExclution struct {
	ExclutionResourceID   string    `gorm:"primaryKey;column:exclution_resource_id" json:"-"`   // 条件值
	ExclutionResourceType bool      `gorm:"primaryKey;column:exclution_resource_type" json:"-"` // 条件类型  1 城市id 2 类目id 3 演出id
	Channel               string    `gorm:"column:channel" json:"channel"`                      // 渠道
	CreateTime            time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime            time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted             []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgSyncExclution) TableName() string {
	return "bg_sync_exclution"
}

// BgSyncExclutionColumns get sql column name.获取数据库列名
var BgSyncExclutionColumns = struct {
	ExclutionResourceID   string
	ExclutionResourceType string
	Channel               string
	CreateTime            string
	UpdateTime            string
	IsDeleted             string
}{
	ExclutionResourceID:   "exclution_resource_id",
	ExclutionResourceType: "exclution_resource_type",
	Channel:               "channel",
	CreateTime:            "create_time",
	UpdateTime:            "update_time",
	IsDeleted:             "is_deleted",
}

// BgSyncSeatplan [...]
type BgSyncSeatplan struct {
	SeatplanID       string    `gorm:"column:seatplan_id" json:"seatplanId"`
	Index            int       `gorm:"column:index" json:"index"`
	ShowTime         time.Time `gorm:"column:show_time" json:"showTime"`
	SyncStatus       []uint8   `gorm:"column:sync_status" json:"syncStatus"` // 数据状态 0：未被同步；1：已同步
	Channel          string    `gorm:"primaryKey;column:channel" json:"-"`   // 渠道编号
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted        []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
	BgSyncSeatplanID string    `gorm:"primaryKey;column:bg_sync_seatplan_id" json:"-"`
}

// TableName get sql table name.获取数据库表名
func (m *BgSyncSeatplan) TableName() string {
	return "bg_sync_seatplan"
}

// BgSyncSeatplanColumns get sql column name.获取数据库列名
var BgSyncSeatplanColumns = struct {
	SeatplanID       string
	Index            string
	ShowTime         string
	SyncStatus       string
	Channel          string
	CreateTime       string
	UpdateTime       string
	IsDeleted        string
	BgSyncSeatplanID string
}{
	SeatplanID:       "seatplan_id",
	Index:            "index",
	ShowTime:         "show_time",
	SyncStatus:       "sync_status",
	Channel:          "channel",
	CreateTime:       "create_time",
	UpdateTime:       "update_time",
	IsDeleted:        "is_deleted",
	BgSyncSeatplanID: "bg_sync_seatplan_id",
}

// BgThrdshowMapping [...]
type BgThrdshowMapping struct {
	ID              int       `gorm:"primaryKey;column:id" json:"-"`
	TkingShowOID    string    `gorm:"column:tkingShowOID" json:"tkingShowOId"`
	TkingShowName   string    `gorm:"column:tkingShowName" json:"tkingShowName"`
	ThrdShowOID     string    `gorm:"column:thrdShowOID" json:"thrdShowOId"`
	FakeThrdShowOID string    `gorm:"column:fakeThrdShowOID" json:"fakeThrdShowOId"`
	ThrdShowName    string    `gorm:"column:thrdShowName" json:"thrdShowName"`
	Source          string    `gorm:"column:source" json:"source"`
	IsDeleted       []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
	CreateTime      time.Time `gorm:"column:createTime" json:"createTime"`
}

// TableName get sql table name.获取数据库表名
func (m *BgThrdshowMapping) TableName() string {
	return "bg_thrdshow_mapping"
}

// BgThrdshowMappingColumns get sql column name.获取数据库列名
var BgThrdshowMappingColumns = struct {
	ID              string
	TkingShowOID    string
	TkingShowName   string
	ThrdShowOID     string
	FakeThrdShowOID string
	ThrdShowName    string
	Source          string
	IsDeleted       string
	CreateTime      string
}{
	ID:              "id",
	TkingShowOID:    "tkingShowOID",
	TkingShowName:   "tkingShowName",
	ThrdShowOID:     "thrdShowOID",
	FakeThrdShowOID: "fakeThrdShowOID",
	ThrdShowName:    "thrdShowName",
	Source:          "source",
	IsDeleted:       "isDeleted",
	CreateTime:      "createTime",
}

// BgTs [...]
type BgTs struct {
	UID  int `gorm:"column:uid" json:"uid"`
	Cofe int `gorm:"column:cofe" json:"cofe"`
	Amt  int `gorm:"column:amt" json:"amt"`
}

// TableName get sql table name.获取数据库表名
func (m *BgTs) TableName() string {
	return "bg_ts"
}

// BgTsColumns get sql column name.获取数据库列名
var BgTsColumns = struct {
	UID  string
	Cofe string
	Amt  string
}{
	UID:  "uid",
	Cofe: "cofe",
	Amt:  "amt",
}

// BgZbfSessionWithSupportexpress [...]
type BgZbfSessionWithSupportexpress struct {
	SessionOID     string    `gorm:"primaryKey;column:sessionOID" json:"-"` // 主办方场次ID
	SupportExpress bool      `gorm:"column:supportExpress" json:"supportExpress"`
	CreateTime     time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime     time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted      bool      `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *BgZbfSessionWithSupportexpress) TableName() string {
	return "bg_zbf_session_with_supportexpress"
}

// BgZbfSessionWithSupportexpressColumns get sql column name.获取数据库列名
var BgZbfSessionWithSupportexpressColumns = struct {
	SessionOID     string
	SupportExpress string
	CreateTime     string
	UpdateTime     string
	IsDeleted      string
}{
	SessionOID:     "sessionOID",
	SupportExpress: "supportExpress",
	CreateTime:     "createTime",
	UpdateTime:     "updateTime",
	IsDeleted:      "isDeleted",
}

// Channel 销售渠道表
type Channel struct {
	ID                 string    `gorm:"primaryKey;column:id" json:"-"`
	BizCode            string    `gorm:"column:biz_code" json:"bizCode"`                        // 业务主体编码
	SubBizCode         string    `gorm:"column:sub_biz_code" json:"subBizCode"`                 // 二级业务编码
	ChannelCategory    string    `gorm:"column:channel_category" json:"channelCategory"`        // 渠道性质取值包括：主站渠道、分销渠道
	ChannelName        string    `gorm:"column:channel_name" json:"channelName"`                // 渠道名称
	ChannelShortName   string    `gorm:"column:channel_short_name" json:"channelShortName"`     // 渠道简称
	ChannelCode        string    `gorm:"column:channel_code" json:"channelCode"`                // 渠道号
	ChannelSource      string    `gorm:"column:channel_source" json:"channelSource"`            // 渠道来源
	AllShowAvailable   bool      `gorm:"column:all_show_available" json:"allShowAvailable"`     // 是否所有的演出都可以投放到该渠道
	AllTicketAvailable bool      `gorm:"column:all_ticket_available" json:"allTicketAvailable"` // 是否所有的库存都可以投放到该渠道
	TechType           string    `gorm:"column:tech_type" json:"techType"`                      // 渠道对接技术类型
	ClientType         string    `gorm:"column:client_type" json:"clientType"`                  // 客户端分类
	ApplicationType    string    `gorm:"column:application_type" json:"applicationType"`        // 应用类型
	MarketingEnabled   []uint8   `gorm:"column:marketing_enabled" json:"marketingEnabled"`      // 是否可对用户营销
	Contacts           string    `gorm:"column:contacts" json:"contacts"`                       // 联系人
	Cellphone          string    `gorm:"column:cellphone" json:"cellphone"`                     // 联系电话
	Email              string    `gorm:"column:email" json:"email"`                             // 邮箱
	StartTime          time.Time `gorm:"column:start_time" json:"startTime"`                    // 有效开始时间
	EndTime            time.Time `gorm:"column:end_time" json:"endTime"`                        // 有效结束时间
	ChannelStatus      string    `gorm:"column:channel_status" json:"channelStatus"`            // 渠道状态
	ChannelGroupID     string    `gorm:"column:channel_group_id" json:"channelGroupId"`         // 渠道分组id
	IsDeleted          []uint8   `gorm:"column:is_deleted" json:"isDeleted"`                    // 是否删除
	CreateTime         time.Time `gorm:"column:create_time" json:"createTime"`                  // 创建时间
	UpdateTime         time.Time `gorm:"column:update_time" json:"updateTime"`                  // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *Channel) TableName() string {
	return "channel"
}

// ChannelColumns get sql column name.获取数据库列名
var ChannelColumns = struct {
	ID                 string
	BizCode            string
	SubBizCode         string
	ChannelCategory    string
	ChannelName        string
	ChannelShortName   string
	ChannelCode        string
	ChannelSource      string
	AllShowAvailable   string
	AllTicketAvailable string
	TechType           string
	ClientType         string
	ApplicationType    string
	MarketingEnabled   string
	Contacts           string
	Cellphone          string
	Email              string
	StartTime          string
	EndTime            string
	ChannelStatus      string
	ChannelGroupID     string
	IsDeleted          string
	CreateTime         string
	UpdateTime         string
}{
	ID:                 "id",
	BizCode:            "biz_code",
	SubBizCode:         "sub_biz_code",
	ChannelCategory:    "channel_category",
	ChannelName:        "channel_name",
	ChannelShortName:   "channel_short_name",
	ChannelCode:        "channel_code",
	ChannelSource:      "channel_source",
	AllShowAvailable:   "all_show_available",
	AllTicketAvailable: "all_ticket_available",
	TechType:           "tech_type",
	ClientType:         "client_type",
	ApplicationType:    "application_type",
	MarketingEnabled:   "marketing_enabled",
	Contacts:           "contacts",
	Cellphone:          "cellphone",
	Email:              "email",
	StartTime:          "start_time",
	EndTime:            "end_time",
	ChannelStatus:      "channel_status",
	ChannelGroupID:     "channel_group_id",
	IsDeleted:          "is_deleted",
	CreateTime:         "create_time",
	UpdateTime:         "update_time",
}

// ChannelComboItemMapping 渠道套票项表
type ChannelComboItemMapping struct {
	ID                       string    `gorm:"primaryKey;column:id" json:"-"`                                       // 主键Id
	ChannelComboItemID       string    `gorm:"column:channel_combo_item_id" json:"channelComboItemId"`              // 第三方渠道套票项Id
	MtcComboItemID           string    `gorm:"column:mtc_combo_item_id" json:"mtcComboItemId"`                      // 中台套票项Id
	ChannelCode              string    `gorm:"column:channel_code" json:"channelCode"`                              // 渠道编码
	ItemPrice                float64   `gorm:"column:item_price" json:"itemPrice"`                                  // 套票项价格
	ItemQty                  int       `gorm:"column:item_qty" json:"itemQty"`                                      // 套票项数量
	ItemStatus               string    `gorm:"column:item_status" json:"itemStatus"`                                // 套票状态：套票项状态：PENDING-待售, ON_SALE-销售中, SOLD_OUT-售罄,HALT-关闭
	SyncStatus               string    `gorm:"column:sync_status" json:"syncStatus"`                                // 同步状态 update_sync_failed-更新同步失败;add_sync_failed-新增同步失败;already_sync-已同步;on_sync-同步中;update_wait_sync-更新待同步;ready_update-更新准备中;add_wait_sync-新增待同步;ready_add-新增准备中
	MtcComboShowID           string    `gorm:"column:mtc_combo_show_id" json:"mtcComboShowId"`                      // 中台套票演出ID
	MtcComboSessionID        string    `gorm:"column:mtc_combo_session_id" json:"mtcComboSessionId"`                // 中台套票场次ID
	MtcComboID               string    `gorm:"column:mtc_combo_id" json:"mtcComboId"`                               // 中台套票ID
	ChannelShowMappingID     string    `gorm:"column:channel_show_mapping_id" json:"channelShowMappingId"`          // 渠道节目ID
	ChannelSessionMappingID  string    `gorm:"column:channel_session_mapping_id" json:"channelSessionMappingId"`    // 渠道场次ID
	ChannelSeatPlanMappingID string    `gorm:"column:channel_seat_plan_mapping_id" json:"channelSeatPlanMappingId"` // 渠道票面ID
	ChannelComboMappingID    string    `gorm:"column:channel_combo_mapping_id" json:"channelComboMappingId"`        // 渠道套票ID
	Extension                string    `gorm:"column:extension" json:"extension"`                                   // 扩展属性，json
	CreateTime               time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime               time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted                bool      `gorm:"column:is_deleted" json:"isDeleted"` // 是否删除
}

// TableName get sql table name.获取数据库表名
func (m *ChannelComboItemMapping) TableName() string {
	return "channel_combo_item_mapping"
}

// ChannelComboItemMappingColumns get sql column name.获取数据库列名
var ChannelComboItemMappingColumns = struct {
	ID                       string
	ChannelComboItemID       string
	MtcComboItemID           string
	ChannelCode              string
	ItemPrice                string
	ItemQty                  string
	ItemStatus               string
	SyncStatus               string
	MtcComboShowID           string
	MtcComboSessionID        string
	MtcComboID               string
	ChannelShowMappingID     string
	ChannelSessionMappingID  string
	ChannelSeatPlanMappingID string
	ChannelComboMappingID    string
	Extension                string
	CreateTime               string
	UpdateTime               string
	IsDeleted                string
}{
	ID:                       "id",
	ChannelComboItemID:       "channel_combo_item_id",
	MtcComboItemID:           "mtc_combo_item_id",
	ChannelCode:              "channel_code",
	ItemPrice:                "item_price",
	ItemQty:                  "item_qty",
	ItemStatus:               "item_status",
	SyncStatus:               "sync_status",
	MtcComboShowID:           "mtc_combo_show_id",
	MtcComboSessionID:        "mtc_combo_session_id",
	MtcComboID:               "mtc_combo_id",
	ChannelShowMappingID:     "channel_show_mapping_id",
	ChannelSessionMappingID:  "channel_session_mapping_id",
	ChannelSeatPlanMappingID: "channel_seat_plan_mapping_id",
	ChannelComboMappingID:    "channel_combo_mapping_id",
	Extension:                "extension",
	CreateTime:               "create_time",
	UpdateTime:               "update_time",
	IsDeleted:                "is_deleted",
}

// ChannelComboMapping 渠道套票表
type ChannelComboMapping struct {
	ID                string    `gorm:"primaryKey;column:id" json:"-"`                        // 主键Id
	ChannelComboID    string    `gorm:"column:channel_combo_id" json:"channelComboId"`        // 第三方渠道套票Id
	MtcComboID        string    `gorm:"column:mtc_combo_id" json:"mtcComboId"`                // 中台套票Id
	ChannelCode       string    `gorm:"column:channel_code" json:"channelCode"`               // 渠道编码
	ComboName         string    `gorm:"column:combo_name" json:"comboName"`                   // 套票名称
	ComboContent      string    `gorm:"column:combo_content" json:"comboContent"`             // 套票说明
	ComboType         string    `gorm:"column:combo_type" json:"comboType"`                   // 套票类型：SAME_SEAT_PLAN-同场次同票面 CROSS_SEAT_PLAN- 同场次跨票面 CROSS_SESSION-跨场次 CROSS_SESSION-跨演出
	ComboPrice        float64   `gorm:"column:combo_price" json:"comboPrice"`                 // 套票价格
	ComboMode         string    `gorm:"column:combo_mode" json:"comboMode"`                   // 套票模式：FREE - 自由套票 FIX- 固定套票
	ComboQty          int       `gorm:"column:combo_qty" json:"comboQty"`                     // 通票定义数量
	ComboUnit         string    `gorm:"column:combo_unit" json:"comboUnit"`                   // 套票单位： TAO-套
	AssignSeatType    string    `gorm:"column:assign_seat_type" json:"assignSeatType"`        // 是否支持选座: NONE_SEAT-非选座 COMMON_SEAT-选座不连坐 LINKED_SEAT- 选座连坐
	StartTime         time.Time `gorm:"column:start_time" json:"startTime"`                   // 套票开始时间
	EndTime           time.Time `gorm:"column:end_time" json:"endTime"`                       // 套票结束时间
	ComboStatus       string    `gorm:"column:combo_status" json:"comboStatus"`               // 套票状态：套票项状态：PENDING-待售, ON_SALE-销售中, SOLD_OUT-售罄,HALT-关闭
	SyncStatus        string    `gorm:"column:sync_status" json:"syncStatus"`                 // 同步状态 update_sync_failed-更新同步失败;add_sync_failed-新增同步失败;already_sync-已同步;on_sync-同步中;update_wait_sync-更新待同步;ready_update-更新准备中;add_wait_sync-新增待同步;ready_add-新增准备中
	MtcComboShowID    string    `gorm:"column:mtc_combo_show_id" json:"mtcComboShowId"`       // 中台套票演出ID
	MtcComboSessionID string    `gorm:"column:mtc_combo_session_id" json:"mtcComboSessionId"` // 中台套票场次ID
	Extension         string    `gorm:"column:extension" json:"extension"`                    // 扩展属性，json
	CreateTime        time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime        time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted         bool      `gorm:"column:is_deleted" json:"isDeleted"` // 是否删除
}

// TableName get sql table name.获取数据库表名
func (m *ChannelComboMapping) TableName() string {
	return "channel_combo_mapping"
}

// ChannelComboMappingColumns get sql column name.获取数据库列名
var ChannelComboMappingColumns = struct {
	ID                string
	ChannelComboID    string
	MtcComboID        string
	ChannelCode       string
	ComboName         string
	ComboContent      string
	ComboType         string
	ComboPrice        string
	ComboMode         string
	ComboQty          string
	ComboUnit         string
	AssignSeatType    string
	StartTime         string
	EndTime           string
	ComboStatus       string
	SyncStatus        string
	MtcComboShowID    string
	MtcComboSessionID string
	Extension         string
	CreateTime        string
	UpdateTime        string
	IsDeleted         string
}{
	ID:                "id",
	ChannelComboID:    "channel_combo_id",
	MtcComboID:        "mtc_combo_id",
	ChannelCode:       "channel_code",
	ComboName:         "combo_name",
	ComboContent:      "combo_content",
	ComboType:         "combo_type",
	ComboPrice:        "combo_price",
	ComboMode:         "combo_mode",
	ComboQty:          "combo_qty",
	ComboUnit:         "combo_unit",
	AssignSeatType:    "assign_seat_type",
	StartTime:         "start_time",
	EndTime:           "end_time",
	ComboStatus:       "combo_status",
	SyncStatus:        "sync_status",
	MtcComboShowID:    "mtc_combo_show_id",
	MtcComboSessionID: "mtc_combo_session_id",
	Extension:         "extension",
	CreateTime:        "create_time",
	UpdateTime:        "update_time",
	IsDeleted:         "is_deleted",
}

// ChannelComboTicketMapping 渠道套票库存表
type ChannelComboTicketMapping struct {
	ID                    string    `gorm:"primaryKey;column:id" json:"-"`                                // 主键Id
	ChannelComboTicketID  string    `gorm:"column:channel_combo_ticket_id" json:"channelComboTicketId"`   // 第三方渠道套票库存Id
	MtcComboTicketID      string    `gorm:"column:mtc_combo_ticket_id" json:"mtcComboTicketId"`           // 中台套票库存Id
	ChannelCode           string    `gorm:"column:channel_code" json:"channelCode"`                       // 渠道编码
	CostPrice             float64   `gorm:"column:cost_price" json:"costPrice"`                           // 套票成本价格
	SalePrice             float64   `gorm:"column:sale_price" json:"salePrice"`                           // 套票销售价格
	LeftStock             int       `gorm:"column:left_stock" json:"leftStock"`                           // 剩余库存
	TicketStatus          string    `gorm:"column:ticket_status" json:"ticketStatus"`                     // 套票项状态：PENDING-待售, ON_SALE-销售中, SOLD_OUT-售罄,HALT-关闭
	SyncStatus            string    `gorm:"column:sync_status" json:"syncStatus"`                         // 同步状态 update_sync_failed-更新同步失败;add_sync_failed-新增同步失败;already_sync-已同步;on_sync-同步中;update_wait_sync-更新待同步;ready_update-更新准备中;add_wait_sync-新增待同步;ready_add-新增准备中
	MtcComboShowID        string    `gorm:"column:mtc_combo_show_id" json:"mtcComboShowId"`               // 中台套票演出ID
	MtcComboSessionID     string    `gorm:"column:mtc_combo_session_id" json:"mtcComboSessionId"`         // 中台套票场次ID
	MtcComboID            string    `gorm:"column:mtc_combo_id" json:"mtcComboId"`                        // 中台套票ID
	ChannelComboMappingID string    `gorm:"column:channel_combo_mapping_id" json:"channelComboMappingId"` // 渠道套票ID
	Extension             string    `gorm:"column:extension" json:"extension"`                            // 扩展属性，json
	CreateTime            time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime            time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted             bool      `gorm:"column:is_deleted" json:"isDeleted"` // 是否删除
}

// TableName get sql table name.获取数据库表名
func (m *ChannelComboTicketMapping) TableName() string {
	return "channel_combo_ticket_mapping"
}

// ChannelComboTicketMappingColumns get sql column name.获取数据库列名
var ChannelComboTicketMappingColumns = struct {
	ID                    string
	ChannelComboTicketID  string
	MtcComboTicketID      string
	ChannelCode           string
	CostPrice             string
	SalePrice             string
	LeftStock             string
	TicketStatus          string
	SyncStatus            string
	MtcComboShowID        string
	MtcComboSessionID     string
	MtcComboID            string
	ChannelComboMappingID string
	Extension             string
	CreateTime            string
	UpdateTime            string
	IsDeleted             string
}{
	ID:                    "id",
	ChannelComboTicketID:  "channel_combo_ticket_id",
	MtcComboTicketID:      "mtc_combo_ticket_id",
	ChannelCode:           "channel_code",
	CostPrice:             "cost_price",
	SalePrice:             "sale_price",
	LeftStock:             "left_stock",
	TicketStatus:          "ticket_status",
	SyncStatus:            "sync_status",
	MtcComboShowID:        "mtc_combo_show_id",
	MtcComboSessionID:     "mtc_combo_session_id",
	MtcComboID:            "mtc_combo_id",
	ChannelComboMappingID: "channel_combo_mapping_id",
	Extension:             "extension",
	CreateTime:            "create_time",
	UpdateTime:            "update_time",
	IsDeleted:             "is_deleted",
}

// ChannelComboTicketSeatMapping 渠道套票库存座位表
type ChannelComboTicketSeatMapping struct {
	ID                          string    `gorm:"primaryKey;column:id" json:"-"`                                             // 主键Id
	ChannelTicketSeatID         string    `gorm:"column:channel_ticket_seat_id" json:"channelTicketSeatId"`                  // 第三方渠道套票库存Id
	MtcTicketSeatID             string    `gorm:"column:mtc_ticket_seat_id" json:"mtcTicketSeatId"`                          // 中台套票库存Id
	ChannelCode                 string    `gorm:"column:channel_code" json:"channelCode"`                                    // 渠道编码
	TicketSeatRule              string    `gorm:"column:ticket_seat_rule" json:"ticketSeatRule"`                             // 座位规则：WHITE_LIST-白名单, BLACK_LIST-黑名单
	TicketSeatType              string    `gorm:"column:ticket_seat_type" json:"ticketSeatType"`                             // 座位类型：SPECIFY_SEAT-指定座位,ZONE_SEAT-区域座位
	TicketSeatStatus            string    `gorm:"column:ticket_seat_status" json:"ticketSeatStatus"`                         // 套票项状态：LOCKING-锁定, ON_SALE-待售, SOLD_OUT-售罄
	SyncStatus                  string    `gorm:"column:sync_status" json:"syncStatus"`                                      // 同步状态 update_sync_failed-更新同步失败;add_sync_failed-新增同步失败;already_sync-已同步;on_sync-同步中;update_wait_sync-更新待同步;ready_update-更新准备中;add_wait_sync-新增待同步;ready_add-新增准备中
	ChannelShowMappingID        string    `gorm:"column:channel_show_mapping_id" json:"channelShowMappingId"`                // 渠道节目ID
	ChannelSessionMappingID     string    `gorm:"column:channel_session_mapping_id" json:"channelSessionMappingId"`          // 渠道场次ID
	ChannelSeatPlanMappingID    string    `gorm:"column:channel_seat_plan_mapping_id" json:"channelSeatPlanMappingId"`       // 渠道票面ID
	ChannelTicketMappingID      string    `gorm:"column:channel_ticket_mapping_id" json:"channelTicketMappingId"`            // 渠道库存ID
	ChannelTicketSeatMappingID  string    `gorm:"column:channel_ticket_seat_mapping_id" json:"channelTicketSeatMappingId"`   // 渠道库存座位ID
	ChannelComboTicketMappingID string    `gorm:"column:channel_combo_ticket_mapping_id" json:"channelComboTicketMappingId"` // 渠道套票库存id
	Extension                   string    `gorm:"column:extension" json:"extension"`                                         // 扩展属性，json
	CreateTime                  time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime                  time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted                   bool      `gorm:"column:is_deleted" json:"isDeleted"` // 是否删除
}

// TableName get sql table name.获取数据库表名
func (m *ChannelComboTicketSeatMapping) TableName() string {
	return "channel_combo_ticket_seat_mapping"
}

// ChannelComboTicketSeatMappingColumns get sql column name.获取数据库列名
var ChannelComboTicketSeatMappingColumns = struct {
	ID                          string
	ChannelTicketSeatID         string
	MtcTicketSeatID             string
	ChannelCode                 string
	TicketSeatRule              string
	TicketSeatType              string
	TicketSeatStatus            string
	SyncStatus                  string
	ChannelShowMappingID        string
	ChannelSessionMappingID     string
	ChannelSeatPlanMappingID    string
	ChannelTicketMappingID      string
	ChannelTicketSeatMappingID  string
	ChannelComboTicketMappingID string
	Extension                   string
	CreateTime                  string
	UpdateTime                  string
	IsDeleted                   string
}{
	ID:                          "id",
	ChannelTicketSeatID:         "channel_ticket_seat_id",
	MtcTicketSeatID:             "mtc_ticket_seat_id",
	ChannelCode:                 "channel_code",
	TicketSeatRule:              "ticket_seat_rule",
	TicketSeatType:              "ticket_seat_type",
	TicketSeatStatus:            "ticket_seat_status",
	SyncStatus:                  "sync_status",
	ChannelShowMappingID:        "channel_show_mapping_id",
	ChannelSessionMappingID:     "channel_session_mapping_id",
	ChannelSeatPlanMappingID:    "channel_seat_plan_mapping_id",
	ChannelTicketMappingID:      "channel_ticket_mapping_id",
	ChannelTicketSeatMappingID:  "channel_ticket_seat_mapping_id",
	ChannelComboTicketMappingID: "channel_combo_ticket_mapping_id",
	Extension:                   "extension",
	CreateTime:                  "create_time",
	UpdateTime:                  "update_time",
	IsDeleted:                   "is_deleted",
}

// ChannelGroup 销售渠道分组表
type ChannelGroup struct {
	ID                string    `gorm:"primaryKey;column:id" json:"-"`
	BizCode           string    `gorm:"column:biz_code" json:"bizCode"`                     // 业务主体编码
	SubBizCode        string    `gorm:"column:sub_biz_code" json:"subBizCode"`              // 二级业务编码
	Name              string    `gorm:"column:name" json:"name"`                            // 渠道分组名称
	Status            string    `gorm:"column:status" json:"status"`                        // 渠道分组状态
	IsDefault         bool      `gorm:"column:is_default" json:"isDefault"`                 // 当前分组是否是商户渠道的默认分组
	DefaultDistribute bool      `gorm:"column:default_distribute" json:"defaultDistribute"` // 当前分组是否默认开启
	IsDeleted         []uint8   `gorm:"column:is_deleted" json:"isDeleted"`                 // 是否删除
	CreateTime        time.Time `gorm:"column:create_time" json:"createTime"`               // 创建时间
	UpdateTime        time.Time `gorm:"column:update_time" json:"updateTime"`               // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ChannelGroup) TableName() string {
	return "channel_group"
}

// ChannelGroupColumns get sql column name.获取数据库列名
var ChannelGroupColumns = struct {
	ID                string
	BizCode           string
	SubBizCode        string
	Name              string
	Status            string
	IsDefault         string
	DefaultDistribute string
	IsDeleted         string
	CreateTime        string
	UpdateTime        string
}{
	ID:                "id",
	BizCode:           "biz_code",
	SubBizCode:        "sub_biz_code",
	Name:              "name",
	Status:            "status",
	IsDefault:         "is_default",
	DefaultDistribute: "default_distribute",
	IsDeleted:         "is_deleted",
	CreateTime:        "create_time",
	UpdateTime:        "update_time",
}

// ChannelOrderCheckCodeInfo 订单核销信息表
type ChannelOrderCheckCodeInfo struct {
	ID             uint64    `gorm:"primaryKey;column:id" json:"-"`
	Channel        string    `gorm:"column:channel" json:"channel"`               // 渠道
	OrderNumber    string    `gorm:"column:orderNumber" json:"orderNumber"`       // 渠道订单号
	OutOrderNumber string    `gorm:"column:outOrderNumber" json:"outOrderNumber"` // 平台订单号
	CheckCode      string    `gorm:"column:checkCode" json:"checkCode"`           // 核销码
	CheckURL       string    `gorm:"column:checkUrl" json:"checkUrl"`             // 核销二维码地址
	CheckMessage   string    `gorm:"column:checkMessage" json:"checkMessage"`     // 核销短信
	SellerID       string    `gorm:"column:sellerId" json:"sellerId"`             // 卖家id
	CreateTime     time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime     time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted      []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *ChannelOrderCheckCodeInfo) TableName() string {
	return "channel_order_check_code_info"
}

// ChannelOrderCheckCodeInfoColumns get sql column name.获取数据库列名
var ChannelOrderCheckCodeInfoColumns = struct {
	ID             string
	Channel        string
	OrderNumber    string
	OutOrderNumber string
	CheckCode      string
	CheckURL       string
	CheckMessage   string
	SellerID       string
	CreateTime     string
	UpdateTime     string
	IsDeleted      string
}{
	ID:             "id",
	Channel:        "channel",
	OrderNumber:    "orderNumber",
	OutOrderNumber: "outOrderNumber",
	CheckCode:      "checkCode",
	CheckURL:       "checkUrl",
	CheckMessage:   "checkMessage",
	SellerID:       "sellerId",
	CreateTime:     "createTime",
	UpdateTime:     "updateTime",
	IsDeleted:      "isDeleted",
}

// ChannelPricingRule 销售渠道定价规则表
type ChannelPricingRule struct {
	ID                      string    `gorm:"primaryKey;column:id" json:"-"`
	ChannelID               string    `gorm:"column:channel_id" json:"channelId"`                               // 销售渠道ID
	PricingType             string    `gorm:"column:pricing_type" json:"pricingType"`                           // 定价方式
	SaleFactorType          string    `gorm:"column:sale_factor_type" json:"saleFactorType"`                    // 销售系数设置类型
	GlobalSaleFactor        float64   `gorm:"column:global_sale_factor" json:"globalSaleFactor"`                // 全品类销售系数值
	IsPriceChange           []uint8   `gorm:"column:is_price_change" json:"isPriceChange"`                      // 是否允许价格变动
	IsFixedSellerStock      []uint8   `gorm:"column:is_fixed_seller_stock" json:"isFixedSellerStock"`           // 是否指定卖家库存
	Flexibility             string    `gorm:"column:flexibility" json:"flexibility"`                            // 定价规则灵活度
	IsPlatServiceFee        bool      `gorm:"column:is_plat_service_fee" json:"isPlatServiceFee"`               // 是否有平台服务费
	PlatServiceRate         float64   `gorm:"column:plat_service_rate" json:"platServiceRate"`                  // 平台服务费
	IsDeleted               []uint8   `gorm:"column:is_deleted" json:"isDeleted"`                               // 是否删除
	CreateTime              time.Time `gorm:"column:create_time" json:"createTime"`                             // 创建时间
	UpdateTime              time.Time `gorm:"column:update_time" json:"updateTime"`                             // 更新时间
	PremiumRatio            float64   `gorm:"column:premium_ratio" json:"premiumRatio"`                         // 随机票溢价比例
	SeatPickingPremiumRatio float64   `gorm:"column:seat_picking_premium_ratio" json:"seatPickingPremiumRatio"` // 选座票溢价比例
}

// TableName get sql table name.获取数据库表名
func (m *ChannelPricingRule) TableName() string {
	return "channel_pricing_rule"
}

// ChannelPricingRuleColumns get sql column name.获取数据库列名
var ChannelPricingRuleColumns = struct {
	ID                      string
	ChannelID               string
	PricingType             string
	SaleFactorType          string
	GlobalSaleFactor        string
	IsPriceChange           string
	IsFixedSellerStock      string
	Flexibility             string
	IsPlatServiceFee        string
	PlatServiceRate         string
	IsDeleted               string
	CreateTime              string
	UpdateTime              string
	PremiumRatio            string
	SeatPickingPremiumRatio string
}{
	ID:                      "id",
	ChannelID:               "channel_id",
	PricingType:             "pricing_type",
	SaleFactorType:          "sale_factor_type",
	GlobalSaleFactor:        "global_sale_factor",
	IsPriceChange:           "is_price_change",
	IsFixedSellerStock:      "is_fixed_seller_stock",
	Flexibility:             "flexibility",
	IsPlatServiceFee:        "is_plat_service_fee",
	PlatServiceRate:         "plat_service_rate",
	IsDeleted:               "is_deleted",
	CreateTime:              "create_time",
	UpdateTime:              "update_time",
	PremiumRatio:            "premium_ratio",
	SeatPickingPremiumRatio: "seat_picking_premium_ratio",
}

// ChannelPricingTemplate 渠道模版
type ChannelPricingTemplate struct {
	ID                string    `gorm:"primaryKey;column:id" json:"-"`                       // 模版id
	TemplateName      string    `gorm:"column:template_name" json:"templateName"`            // 渠道名称
	VocalConcertRatio float64   `gorm:"column:vocal_concert_ratio" json:"vocalConcertRatio"` // 演唱会系数
	ConcertRatio      float64   `gorm:"column:concert_ratio" json:"concertRatio"`            // 音乐会系数
	DramaRatio        float64   `gorm:"column:drama_ratio" json:"dramaRatio"`                // 话剧歌剧系数
	AcrobaticsRatio   float64   `gorm:"column:acrobatics_ratio" json:"acrobaticsRatio"`      // 曲艺杂谈系数
	DancingRatio      float64   `gorm:"column:dancing_ratio" json:"dancingRatio"`            // 舞蹈芭蕾系数
	SportsRatio       float64   `gorm:"column:sports_ratio" json:"sportsRatio"`              // 体育赛事系数
	ExhibitionRatio   float64   `gorm:"column:exhibition_ratio" json:"exhibitionRatio"`      // 展览休闲系数
	ChildrenRatio     float64   `gorm:"column:children_ratio" json:"childrenRatio"`          // 儿童亲子系数
	IsDeleted         []uint8   `gorm:"column:is_deleted" json:"isDeleted"`                  // 是否删除
	CreateTime        time.Time `gorm:"column:create_time" json:"createTime"`                // 创建时间
	UpdateTime        time.Time `gorm:"column:update_time" json:"updateTime"`                // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ChannelPricingTemplate) TableName() string {
	return "channel_pricing_template"
}

// ChannelPricingTemplateColumns get sql column name.获取数据库列名
var ChannelPricingTemplateColumns = struct {
	ID                string
	TemplateName      string
	VocalConcertRatio string
	ConcertRatio      string
	DramaRatio        string
	AcrobaticsRatio   string
	DancingRatio      string
	SportsRatio       string
	ExhibitionRatio   string
	ChildrenRatio     string
	IsDeleted         string
	CreateTime        string
	UpdateTime        string
}{
	ID:                "id",
	TemplateName:      "template_name",
	VocalConcertRatio: "vocal_concert_ratio",
	ConcertRatio:      "concert_ratio",
	DramaRatio:        "drama_ratio",
	AcrobaticsRatio:   "acrobatics_ratio",
	DancingRatio:      "dancing_ratio",
	SportsRatio:       "sports_ratio",
	ExhibitionRatio:   "exhibition_ratio",
	ChildrenRatio:     "children_ratio",
	IsDeleted:         "is_deleted",
	CreateTime:        "create_time",
	UpdateTime:        "update_time",
}

// ChannelSeatPlanColor 票面价格颜色表
type ChannelSeatPlanColor struct {
	ID                   string    `gorm:"primaryKey;column:id" json:"-"`                              // 主键
	ChannelShowMappingID string    `gorm:"column:channel_show_mapping_id" json:"channelShowMappingId"` // 演出映射Id
	SeatPlanPrice        float64   `gorm:"column:seat_plan_price" json:"seatPlanPrice"`                // 售价
	ChannelCode          string    `gorm:"column:channel_code" json:"channelCode"`                     // 渠道编码
	Color                string    `gorm:"column:color" json:"color"`                                  // 颜色值
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`                       // 创建时刻
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`                       // 修改时刻
	IsDeleted            bool      `gorm:"column:is_deleted" json:"isDeleted"`                         // 删除标志
}

// TableName get sql table name.获取数据库表名
func (m *ChannelSeatPlanColor) TableName() string {
	return "channel_seat_plan_color"
}

// ChannelSeatPlanColorColumns get sql column name.获取数据库列名
var ChannelSeatPlanColorColumns = struct {
	ID                   string
	ChannelShowMappingID string
	SeatPlanPrice        string
	ChannelCode          string
	Color                string
	CreateTime           string
	UpdateTime           string
	IsDeleted            string
}{
	ID:                   "id",
	ChannelShowMappingID: "channel_show_mapping_id",
	SeatPlanPrice:        "seat_plan_price",
	ChannelCode:          "channel_code",
	Color:                "color",
	CreateTime:           "create_time",
	UpdateTime:           "update_time",
	IsDeleted:            "is_deleted",
}

// ChannelSettleRule 销售渠道结算规则表
type ChannelSettleRule struct {
	ID              string    `gorm:"primaryKey;column:id" json:"-"`
	ChannelID       string    `gorm:"column:channel_id" json:"channelId"`               // 销售渠道ID
	SettleType      string    `gorm:"column:settle_type" json:"settleType"`             // 结算方式
	SettleTimeType  string    `gorm:"column:settle_time_type" json:"settleTimeType"`    // 结算规则
	SettleInterval  int       `gorm:"column:settle_interval" json:"settleInterval"`     // 结算周期（天）
	AgentFeeType    string    `gorm:"column:agent_fee_type" json:"agentFeeType"`        // 代理服务费类型
	AgentFeeGmvRate float64   `gorm:"column:agent_fee_gmv_rate" json:"agentFeeGmvRate"` // 整体销售额方式下的代理服务费
	IsDeleted       []uint8   `gorm:"column:is_deleted" json:"isDeleted"`               // 是否删除
	CreateTime      time.Time `gorm:"column:create_time" json:"createTime"`             // 创建时间
	UpdateTime      time.Time `gorm:"column:update_time" json:"updateTime"`             // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ChannelSettleRule) TableName() string {
	return "channel_settle_rule"
}

// ChannelSettleRuleColumns get sql column name.获取数据库列名
var ChannelSettleRuleColumns = struct {
	ID              string
	ChannelID       string
	SettleType      string
	SettleTimeType  string
	SettleInterval  string
	AgentFeeType    string
	AgentFeeGmvRate string
	IsDeleted       string
	CreateTime      string
	UpdateTime      string
}{
	ID:              "id",
	ChannelID:       "channel_id",
	SettleType:      "settle_type",
	SettleTimeType:  "settle_time_type",
	SettleInterval:  "settle_interval",
	AgentFeeType:    "agent_fee_type",
	AgentFeeGmvRate: "agent_fee_gmv_rate",
	IsDeleted:       "is_deleted",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
}

// ChannelShowBlackList 销售渠道演出投放黑名单表
type ChannelShowBlackList struct {
	ID         string    `gorm:"primaryKey;column:id" json:"-"`
	ChannelID  string    `gorm:"column:channel_id" json:"channelId"`   // 销售渠道ID
	ShowID     string    `gorm:"column:show_id" json:"showId"`         // 演出ID
	ShowName   string    `gorm:"column:show_name" json:"showName"`     // 演出名称
	IsDeleted  []uint8   `gorm:"column:is_deleted" json:"isDeleted"`   // 是否删除
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ChannelShowBlackList) TableName() string {
	return "channel_show_black_list"
}

// ChannelShowBlackListColumns get sql column name.获取数据库列名
var ChannelShowBlackListColumns = struct {
	ID         string
	ChannelID  string
	ShowID     string
	ShowName   string
	IsDeleted  string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	ChannelID:  "channel_id",
	ShowID:     "show_id",
	ShowName:   "show_name",
	IsDeleted:  "is_deleted",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// ChannelShowDeliverMethodRule 渠道项目配送方式
type ChannelShowDeliverMethodRule struct {
	ID            string    `gorm:"primaryKey;column:id" json:"-"`               // 主键ID
	ChannelID     string    `gorm:"column:channel_id" json:"channelId"`          // 渠道id
	ShowID        string    `gorm:"column:show_id" json:"showId"`                // 项目id
	SourceID      string    `gorm:"column:source_id" json:"sourceId"`            // 规则来源ID，例如：现场购ID
	ShowSessionID string    `gorm:"column:show_session_id" json:"showSessionId"` // 场次id
	DeliverMethod string    `gorm:"column:deliver_method" json:"deliverMethod"`  // 配送方式
	Valid         bool      `gorm:"column:valid" json:"valid"`                   // 是否有效
	IsDeleted     bool      `gorm:"column:is_deleted" json:"isDeleted"`          // 是否删除
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`        // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`        // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ChannelShowDeliverMethodRule) TableName() string {
	return "channel_show_deliver_method_rule"
}

// ChannelShowDeliverMethodRuleColumns get sql column name.获取数据库列名
var ChannelShowDeliverMethodRuleColumns = struct {
	ID            string
	ChannelID     string
	ShowID        string
	SourceID      string
	ShowSessionID string
	DeliverMethod string
	Valid         string
	IsDeleted     string
	CreateTime    string
	UpdateTime    string
}{
	ID:            "id",
	ChannelID:     "channel_id",
	ShowID:        "show_id",
	SourceID:      "source_id",
	ShowSessionID: "show_session_id",
	DeliverMethod: "deliver_method",
	Valid:         "valid",
	IsDeleted:     "is_deleted",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// ChannelShowPricingRule 销售渠道演出级别定价规则表
type ChannelShowPricingRule struct {
	ID                      string    `gorm:"primaryKey;column:id" json:"-"`
	ChannelID               string    `gorm:"column:channel_id" json:"channelId"`                               // 销售渠道ID
	Category                string    `gorm:"column:category" json:"category"`                                  // 演出类目
	ShowID                  string    `gorm:"column:show_id" json:"showId"`                                     // 演出ID
	ShowName                string    `gorm:"column:show_name" json:"showName"`                                 // 演出名称
	SaleFactor              float64   `gorm:"column:sale_factor" json:"saleFactor"`                             // 销售系数
	SupportEditSaleFactor   []uint8   `gorm:"column:support_edit_sale_factor" json:"supportEditSaleFactor"`     // 是否支持动态调整销售系数
	MagnetSaleFactor        float64   `gorm:"column:magnet_sale_factor" json:"magnetSaleFactor"`                // 吸铁石系数
	BlackListFactor         float64   `gorm:"column:black_list_factor" json:"blackListFactor"`                  // 黑名单溢价系数
	IsDeleted               []uint8   `gorm:"column:is_deleted" json:"isDeleted"`                               // 是否删除
	CreateTime              time.Time `gorm:"column:create_time" json:"createTime"`                             // 创建时间
	UpdateTime              time.Time `gorm:"column:update_time" json:"updateTime"`                             // 更新时间
	PremiumRatio            float64   `gorm:"column:premium_ratio" json:"premiumRatio"`                         // 随机票溢价比例
	SeatPickingPremiumRatio float64   `gorm:"column:seat_picking_premium_ratio" json:"seatPickingPremiumRatio"` // 选座票溢价比例
}

// TableName get sql table name.获取数据库表名
func (m *ChannelShowPricingRule) TableName() string {
	return "channel_show_pricing_rule"
}

// ChannelShowPricingRuleColumns get sql column name.获取数据库列名
var ChannelShowPricingRuleColumns = struct {
	ID                      string
	ChannelID               string
	Category                string
	ShowID                  string
	ShowName                string
	SaleFactor              string
	SupportEditSaleFactor   string
	MagnetSaleFactor        string
	BlackListFactor         string
	IsDeleted               string
	CreateTime              string
	UpdateTime              string
	PremiumRatio            string
	SeatPickingPremiumRatio string
}{
	ID:                      "id",
	ChannelID:               "channel_id",
	Category:                "category",
	ShowID:                  "show_id",
	ShowName:                "show_name",
	SaleFactor:              "sale_factor",
	SupportEditSaleFactor:   "support_edit_sale_factor",
	MagnetSaleFactor:        "magnet_sale_factor",
	BlackListFactor:         "black_list_factor",
	IsDeleted:               "is_deleted",
	CreateTime:              "create_time",
	UpdateTime:              "update_time",
	PremiumRatio:            "premium_ratio",
	SeatPickingPremiumRatio: "seat_picking_premium_ratio",
}

// ChannelShowRule 销售渠道演出投放规则表
type ChannelShowRule struct {
	ID                     string    `gorm:"primaryKey;column:id" json:"-"`
	ChannelID              string    `gorm:"column:channel_id" json:"channelId"`                             // 销售渠道ID
	SupportExpress         []uint8   `gorm:"column:support_express" json:"supportExpress"`                   // 是否支持快递送票
	SupportVenue           []uint8   `gorm:"column:support_venue" json:"supportVenue"`                       // 是否只是现场取票
	SupportOnsite          []uint8   `gorm:"column:support_onsite" json:"supportOnsite"`                     // 是否支持上门取票
	SupportSeatPicking     string    `gorm:"column:support_seat_picking" json:"supportSeatPicking"`          // 是否支持选座
	SeatPickingType        string    `gorm:"column:seat_picking_type" json:"seatPickingType"`                // 选座方式：选座到排/选座到座
	SupportRealName        string    `gorm:"column:support_real_name" json:"supportRealName"`                // 是否支持实名
	CardSuggestType        string    `gorm:"column:card_suggest_type" json:"cardSuggestType"`                // 身份证支持类型
	Statuses               string    `gorm:"column:statuses" json:"statuses"`                                // 演出状态
	Status                 string    `gorm:"column:status" json:"status"`                                    // 演出状态
	RealNameType           string    `gorm:"column:real_name_type" json:"realNameType"`                      // 支持的实名类型
	SupportEarlyBirdTicket string    `gorm:"column:support_early_bird_ticket" json:"supportEarlyBirdTicket"` // 是否支持早鸟票
	SupportTailTicket      string    `gorm:"column:support_tail_ticket" json:"supportTailTicket"`            // 是否支持尾票
	Sites                  string    `gorm:"column:sites" json:"sites"`                                      // 包含的演出站点列表
	Categorys              string    `gorm:"column:categorys" json:"categorys"`                              // 包含的演出类目列表
	SupportCommonShow      []uint8   `gorm:"column:support_common_show" json:"supportCommonShow"`            // 是否包含普通演出
	SupportActivityShow    []uint8   `gorm:"column:support_activity_show" json:"supportActivityShow"`        // 是否包含活动演出
	SupportSensitiveShow   []uint8   `gorm:"column:support_sensitive_show" json:"supportSensitiveShow"`      // 是否包含敏感演出
	IsDeleted              []uint8   `gorm:"column:is_deleted" json:"isDeleted"`                             // 是否删除
	CreateTime             time.Time `gorm:"column:create_time" json:"createTime"`                           // 创建时间
	UpdateTime             time.Time `gorm:"column:update_time" json:"updateTime"`                           // 更新时间
	SvgEmpty               bool      `gorm:"column:svg_empty" json:"svgEmpty"`                               // 是否需要svg
	SensitiveAutoBlack     bool      `gorm:"column:sensitive_auto_black" json:"sensitiveAutoBlack"`
}

// TableName get sql table name.获取数据库表名
func (m *ChannelShowRule) TableName() string {
	return "channel_show_rule"
}

// ChannelShowRuleColumns get sql column name.获取数据库列名
var ChannelShowRuleColumns = struct {
	ID                     string
	ChannelID              string
	SupportExpress         string
	SupportVenue           string
	SupportOnsite          string
	SupportSeatPicking     string
	SeatPickingType        string
	SupportRealName        string
	CardSuggestType        string
	Statuses               string
	Status                 string
	RealNameType           string
	SupportEarlyBirdTicket string
	SupportTailTicket      string
	Sites                  string
	Categorys              string
	SupportCommonShow      string
	SupportActivityShow    string
	SupportSensitiveShow   string
	IsDeleted              string
	CreateTime             string
	UpdateTime             string
	SvgEmpty               string
	SensitiveAutoBlack     string
}{
	ID:                     "id",
	ChannelID:              "channel_id",
	SupportExpress:         "support_express",
	SupportVenue:           "support_venue",
	SupportOnsite:          "support_onsite",
	SupportSeatPicking:     "support_seat_picking",
	SeatPickingType:        "seat_picking_type",
	SupportRealName:        "support_real_name",
	CardSuggestType:        "card_suggest_type",
	Statuses:               "statuses",
	Status:                 "status",
	RealNameType:           "real_name_type",
	SupportEarlyBirdTicket: "support_early_bird_ticket",
	SupportTailTicket:      "support_tail_ticket",
	Sites:                  "sites",
	Categorys:              "categorys",
	SupportCommonShow:      "support_common_show",
	SupportActivityShow:    "support_activity_show",
	SupportSensitiveShow:   "support_sensitive_show",
	IsDeleted:              "is_deleted",
	CreateTime:             "create_time",
	UpdateTime:             "update_time",
	SvgEmpty:               "svg_empty",
	SensitiveAutoBlack:     "sensitive_auto_black",
}

// ChannelShowSeatPlanRule 渠道项目票面
type ChannelShowSeatPlanRule struct {
	ID            string    `gorm:"primaryKey;column:id" json:"-"`               // 主键ID
	ChannelID     string    `gorm:"column:channel_id" json:"channelId"`          // 渠道id
	ShowID        string    `gorm:"column:show_id" json:"showId"`                // 项目id
	SourceID      string    `gorm:"column:source_id" json:"sourceId"`            // 规则来源ID，例如：现场购ID
	ShowSessionID string    `gorm:"column:show_session_id" json:"showSessionId"` // 项目场次id
	SeatPlanID    string    `gorm:"column:seat_plan_id" json:"seatPlanId"`       // 项目票面id
	Valid         bool      `gorm:"column:valid" json:"valid"`                   // 是否有效
	IsDeleted     bool      `gorm:"column:is_deleted" json:"isDeleted"`          // 是否删除
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`        // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`        // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ChannelShowSeatPlanRule) TableName() string {
	return "channel_show_seat_plan_rule"
}

// ChannelShowSeatPlanRuleColumns get sql column name.获取数据库列名
var ChannelShowSeatPlanRuleColumns = struct {
	ID            string
	ChannelID     string
	ShowID        string
	SourceID      string
	ShowSessionID string
	SeatPlanID    string
	Valid         string
	IsDeleted     string
	CreateTime    string
	UpdateTime    string
}{
	ID:            "id",
	ChannelID:     "channel_id",
	ShowID:        "show_id",
	SourceID:      "source_id",
	ShowSessionID: "show_session_id",
	SeatPlanID:    "seat_plan_id",
	Valid:         "valid",
	IsDeleted:     "is_deleted",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// ChannelShowSessionRule 渠道项目场次
type ChannelShowSessionRule struct {
	ID            string    `gorm:"primaryKey;column:id" json:"-"`               // 主键ID
	ChannelID     string    `gorm:"column:channel_id" json:"channelId"`          // 渠道id
	ShowID        string    `gorm:"column:show_id" json:"showId"`                // 项目id
	SourceID      string    `gorm:"column:source_id" json:"sourceId"`            // 规则来源ID，例如：现场购ID
	ShowSessionID string    `gorm:"column:show_session_id" json:"showSessionId"` // 项目场次id
	Valid         bool      `gorm:"column:valid" json:"valid"`                   // 是否有效
	IsDeleted     bool      `gorm:"column:is_deleted" json:"isDeleted"`          // 是否删除
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`        // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`        // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ChannelShowSessionRule) TableName() string {
	return "channel_show_session_rule"
}

// ChannelShowSessionRuleColumns get sql column name.获取数据库列名
var ChannelShowSessionRuleColumns = struct {
	ID            string
	ChannelID     string
	ShowID        string
	SourceID      string
	ShowSessionID string
	Valid         string
	IsDeleted     string
	CreateTime    string
	UpdateTime    string
}{
	ID:            "id",
	ChannelID:     "channel_id",
	ShowID:        "show_id",
	SourceID:      "source_id",
	ShowSessionID: "show_session_id",
	Valid:         "valid",
	IsDeleted:     "is_deleted",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// ChannelShowSettleRule 销售渠道演出级别结算规则表
type ChannelShowSettleRule struct {
	ID           string    `gorm:"primaryKey;column:id" json:"-"`
	ChannelID    string    `gorm:"column:channel_id" json:"channelId"`        // 销售渠道ID
	Category     string    `gorm:"column:category" json:"category"`           // 演出类目
	ShowID       string    `gorm:"column:show_id" json:"showId"`              // 演出ID
	ShowName     string    `gorm:"column:show_name" json:"showName"`          // 演出名称
	AgentFeeRate float64   `gorm:"column:agent_fee_rate" json:"agentFeeRate"` // 代理服务费率
	IsDeleted    []uint8   `gorm:"column:is_deleted" json:"isDeleted"`        // 是否删除
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`      // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`      // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ChannelShowSettleRule) TableName() string {
	return "channel_show_settle_rule"
}

// ChannelShowSettleRuleColumns get sql column name.获取数据库列名
var ChannelShowSettleRuleColumns = struct {
	ID           string
	ChannelID    string
	Category     string
	ShowID       string
	ShowName     string
	AgentFeeRate string
	IsDeleted    string
	CreateTime   string
	UpdateTime   string
}{
	ID:           "id",
	ChannelID:    "channel_id",
	Category:     "category",
	ShowID:       "show_id",
	ShowName:     "show_name",
	AgentFeeRate: "agent_fee_rate",
	IsDeleted:    "is_deleted",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// ChannelShowTenantMapping 渠道演出和主办映射表
type ChannelShowTenantMapping struct {
	ID              string    `gorm:"column:id" json:"id"`                             // 映射Id
	ChannelShowID   string    `gorm:"column:channel_show_id" json:"channelShowId"`     // 渠道演出Id
	ChannelTenantID string    `gorm:"column:channel_tenant_id" json:"channelTenantId"` // 渠道主办租户Id
	CreateTime      time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted       bool      `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *ChannelShowTenantMapping) TableName() string {
	return "channel_show_tenant_mapping"
}

// ChannelShowTenantMappingColumns get sql column name.获取数据库列名
var ChannelShowTenantMappingColumns = struct {
	ID              string
	ChannelShowID   string
	ChannelTenantID string
	CreateTime      string
	UpdateTime      string
	IsDeleted       string
}{
	ID:              "id",
	ChannelShowID:   "channel_show_id",
	ChannelTenantID: "channel_tenant_id",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
	IsDeleted:       "is_deleted",
}

// ChannelTicketRule 销售渠道库存投放规则表
type ChannelTicketRule struct {
	ID                     string    `gorm:"primaryKey;column:id" json:"-"`
	ChannelID              string    `gorm:"column:channel_id" json:"channelId"`                // 销售渠道ID
	IsLimitedStock         []uint8   `gorm:"column:is_limited_stock" json:"isLimitedStock"`     // 是否限制库存
	LimitedStockRate       float64   `gorm:"column:limited_stock_rate" json:"limitedStockRate"` // 限制库存比例（占总库存的比例）
	IsDeleted              []uint8   `gorm:"column:is_deleted" json:"isDeleted"`                // 是否删除
	CreateTime             time.Time `gorm:"column:create_time" json:"createTime"`              // 创建时间
	UpdateTime             time.Time `gorm:"column:update_time" json:"updateTime"`              // 更新时间
	OptimizeTicketType     string    `gorm:"column:optimize_ticket_type" json:"optimizeTicketType"`
	ChooseTicket           bool      `gorm:"column:choose_ticket" json:"chooseTicket"`                       // 混合只给纸质票
	DisplayItem            bool      `gorm:"column:display_item" json:"displayItem"`                         // 隐藏字段 costPrice隐藏
	SupportSeatPicking     bool      `gorm:"column:support_seat_picking" json:"supportSeatPicking"`          // 选座票
	SupportQuickDeliver    bool      `gorm:"column:support_quick_deliver" json:"supportQuickDeliver"`        // 极速发货
	SupportETicket         bool      `gorm:"column:support_e_ticket" json:"supportETicket"`                  // 电子票库存
	IsFixedSellerStock     bool      `gorm:"column:is_fixed_seller_stock" json:"isFixedSellerStock"`         // 是否指定卖家库存
	StockType              string    `gorm:"column:stock_type" json:"stockType"`                             // 受限库存类型
	SensitiveShowStockType string    `gorm:"column:sensitive_show_stock_type" json:"sensitiveShowStockType"` // 敏感演出库存类型
}

// TableName get sql table name.获取数据库表名
func (m *ChannelTicketRule) TableName() string {
	return "channel_ticket_rule"
}

// ChannelTicketRuleColumns get sql column name.获取数据库列名
var ChannelTicketRuleColumns = struct {
	ID                     string
	ChannelID              string
	IsLimitedStock         string
	LimitedStockRate       string
	IsDeleted              string
	CreateTime             string
	UpdateTime             string
	OptimizeTicketType     string
	ChooseTicket           string
	DisplayItem            string
	SupportSeatPicking     string
	SupportQuickDeliver    string
	SupportETicket         string
	IsFixedSellerStock     string
	StockType              string
	SensitiveShowStockType string
}{
	ID:                     "id",
	ChannelID:              "channel_id",
	IsLimitedStock:         "is_limited_stock",
	LimitedStockRate:       "limited_stock_rate",
	IsDeleted:              "is_deleted",
	CreateTime:             "create_time",
	UpdateTime:             "update_time",
	OptimizeTicketType:     "optimize_ticket_type",
	ChooseTicket:           "choose_ticket",
	DisplayItem:            "display_item",
	SupportSeatPicking:     "support_seat_picking",
	SupportQuickDeliver:    "support_quick_deliver",
	SupportETicket:         "support_e_ticket",
	IsFixedSellerStock:     "is_fixed_seller_stock",
	StockType:              "stock_type",
	SensitiveShowStockType: "sensitive_show_stock_type",
}

// ChannelTicketSellerRule 销售渠道卖家库存投放规则表
type ChannelTicketSellerRule struct {
	ID               string    `gorm:"primaryKey;column:id" json:"-"`
	ChannelID        string    `gorm:"column:channel_id" json:"channelId"`                // 销售渠道ID
	SellerType       string    `gorm:"column:seller_type" json:"sellerType"`              // 卖家类型
	SellerID         string    `gorm:"column:seller_id" json:"sellerId"`                  // 可投放的卖家ID，表示具体的卖家库存可以被投放
	SellerName       string    `gorm:"column:seller_name" json:"sellerName"`              // 卖家名称
	IsLimitedStock   bool      `gorm:"column:is_limited_stock" json:"isLimitedStock"`     // 是否限制库存
	LimitedStockRate float64   `gorm:"column:limited_stock_rate" json:"limitedStockRate"` // 限制库存比例（占总库存的比例）
	IsDeleted        bool      `gorm:"column:is_deleted" json:"isDeleted"`                // 是否删除
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`              // 创建时间
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`              // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ChannelTicketSellerRule) TableName() string {
	return "channel_ticket_seller_rule"
}

// ChannelTicketSellerRuleColumns get sql column name.获取数据库列名
var ChannelTicketSellerRuleColumns = struct {
	ID               string
	ChannelID        string
	SellerType       string
	SellerID         string
	SellerName       string
	IsLimitedStock   string
	LimitedStockRate string
	IsDeleted        string
	CreateTime       string
	UpdateTime       string
}{
	ID:               "id",
	ChannelID:        "channel_id",
	SellerType:       "seller_type",
	SellerID:         "seller_id",
	SellerName:       "seller_name",
	IsLimitedStock:   "is_limited_stock",
	LimitedStockRate: "limited_stock_rate",
	IsDeleted:        "is_deleted",
	CreateTime:       "create_time",
	UpdateTime:       "update_time",
}

// ChannelTradeRule 销售渠道交易规则表
type ChannelTradeRule struct {
	ID                string    `gorm:"primaryKey;column:id" json:"-"`
	ChannelID         string    `gorm:"column:channel_id" json:"channelId"`                  // 销售渠道ID
	IsSplitOrder      []uint8   `gorm:"column:is_split_order" json:"isSplitOrder"`           // 是否需要拆单
	ExpressFeeType    string    `gorm:"column:express_fee_type" json:"expressFeeType"`       // 运费类型
	ExpressFee        float64   `gorm:"column:express_fee" json:"expressFee"`                // 主快递公司快递费
	OtherExpressFee   float64   `gorm:"column:other_express_fee" json:"otherExpressFee"`     // 其他快递公司快递费
	IsPlatServiceFee  bool      `gorm:"column:is_plat_service_fee" json:"isPlatServiceFee"`  // 是否有平台服务费
	PlatServiceRate   float64   `gorm:"column:plat_service_rate" json:"platServiceRate"`     // 平台服务费率
	LimitationRule    string    `gorm:"column:limitation_rule" json:"limitationRule"`        // 限购规则
	PaymentTypes      string    `gorm:"column:payment_types" json:"paymentTypes"`            // 支付方式
	ShowBrand         string    `gorm:"column:show_brand" json:"showBrand"`                  // 品牌露出
	SmsSendRule       string    `gorm:"column:sms_send_rule" json:"smsSendRule"`             // 短信发送规则
	DeliverTemplateID string    `gorm:"column:deliver_template_id" json:"deliverTemplateId"` // 运费模板ID
	IsDeleted         []uint8   `gorm:"column:is_deleted" json:"isDeleted"`                  // 是否删除
	CreateTime        time.Time `gorm:"column:create_time" json:"createTime"`                // 创建时间
	UpdateTime        time.Time `gorm:"column:update_time" json:"updateTime"`                // 更新时间
	SmsType           string    `gorm:"column:sms_type" json:"smsType"`                      // 短信类型：下单短信(SUBMIT_ORDER)、支付提示(PAY_ORDER)、配送短信(DELIVER_GOODS)、订单取消(CANCLE_ORDER) 可以多选 字符串拼接 逗号分隔
	NotifyChannelID   string    `gorm:"column:notify_channel_id" json:"notifyChannelId"`     // 通知中心的productchannelid
	IsSmsShowChannel  []uint8   `gorm:"column:is_sms_show_channel" json:"isSmsShowChannel"`  // 通知短信中是否显示渠道名称
	QrCodeLinkURL     string    `gorm:"column:qr_code_link_url" json:"qrCodeLinkUrl"`        // 取票链接域名
}

// TableName get sql table name.获取数据库表名
func (m *ChannelTradeRule) TableName() string {
	return "channel_trade_rule"
}

// ChannelTradeRuleColumns get sql column name.获取数据库列名
var ChannelTradeRuleColumns = struct {
	ID                string
	ChannelID         string
	IsSplitOrder      string
	ExpressFeeType    string
	ExpressFee        string
	OtherExpressFee   string
	IsPlatServiceFee  string
	PlatServiceRate   string
	LimitationRule    string
	PaymentTypes      string
	ShowBrand         string
	SmsSendRule       string
	DeliverTemplateID string
	IsDeleted         string
	CreateTime        string
	UpdateTime        string
	SmsType           string
	NotifyChannelID   string
	IsSmsShowChannel  string
	QrCodeLinkURL     string
}{
	ID:                "id",
	ChannelID:         "channel_id",
	IsSplitOrder:      "is_split_order",
	ExpressFeeType:    "express_fee_type",
	ExpressFee:        "express_fee",
	OtherExpressFee:   "other_express_fee",
	IsPlatServiceFee:  "is_plat_service_fee",
	PlatServiceRate:   "plat_service_rate",
	LimitationRule:    "limitation_rule",
	PaymentTypes:      "payment_types",
	ShowBrand:         "show_brand",
	SmsSendRule:       "sms_send_rule",
	DeliverTemplateID: "deliver_template_id",
	IsDeleted:         "is_deleted",
	CreateTime:        "create_time",
	UpdateTime:        "update_time",
	SmsType:           "sms_type",
	NotifyChannelID:   "notify_channel_id",
	IsSmsShowChannel:  "is_sms_show_channel",
	QrCodeLinkURL:     "qr_code_link_url",
}

// ChannelUserMapping 渠道用户表
type ChannelUserMapping struct {
	ID            string    `gorm:"primaryKey;column:id" json:"-"`               // 主键Id
	ChannelUserID string    `gorm:"column:channel_user_id" json:"channelUserId"` // 第三方用户Id
	MtcUserID     string    `gorm:"column:mtc_user_id" json:"mtcUserId"`         // 中台用户Id
	BizUserID     string    `gorm:"column:biz_user_id" json:"bizUserId"`         // 端用户Id
	ChannelCode   string    `gorm:"column:channel_code" json:"channelCode"`      // 渠道编码
	Phone         string    `gorm:"column:phone" json:"phone"`                   // 用户手机号
	UserRealName  string    `gorm:"column:user_real_name" json:"userRealName"`   // 用户真实名称
	UserStatus    string    `gorm:"column:user_status" json:"userStatus"`        // 用户状态：NORMAL-正常, DISABLED-已注销
	SyncStatus    string    `gorm:"column:sync_status" json:"syncStatus"`        // 同步状态 update_sync_failed-更新同步失败;add_sync_failed-新增同步失败;already_sync-已同步;on_sync-同步中;update_wait_sync-更新待同步;ready_update-更新准备中;add_wait_sync-新增待同步;ready_add-新增准备中
	Extension     string    `gorm:"column:extension" json:"extension"`           // 扩展属性，json
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted     bool      `gorm:"column:is_deleted" json:"isDeleted"` // 是否删除
}

// TableName get sql table name.获取数据库表名
func (m *ChannelUserMapping) TableName() string {
	return "channel_user_mapping"
}

// ChannelUserMappingColumns get sql column name.获取数据库列名
var ChannelUserMappingColumns = struct {
	ID            string
	ChannelUserID string
	MtcUserID     string
	BizUserID     string
	ChannelCode   string
	Phone         string
	UserRealName  string
	UserStatus    string
	SyncStatus    string
	Extension     string
	CreateTime    string
	UpdateTime    string
	IsDeleted     string
}{
	ID:            "id",
	ChannelUserID: "channel_user_id",
	MtcUserID:     "mtc_user_id",
	BizUserID:     "biz_user_id",
	ChannelCode:   "channel_code",
	Phone:         "phone",
	UserRealName:  "user_real_name",
	UserStatus:    "user_status",
	SyncStatus:    "sync_status",
	Extension:     "extension",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
	IsDeleted:     "is_deleted",
}

// ChannelVerificationInfo 订单核销信息表
type ChannelVerificationInfo struct {
	ID             uint64    `gorm:"primaryKey;column:id" json:"-"`
	Channel        string    `gorm:"column:channel" json:"channel"`               // 渠道
	OrderNumber    string    `gorm:"column:orderNumber" json:"orderNumber"`       // 渠道订单号
	OutOrderNumber string    `gorm:"column:outOrderNumber" json:"outOrderNumber"` // 平台订单号
	CheckCode      string    `gorm:"column:checkCode" json:"checkCode"`           // 核销码
	CheckURL       string    `gorm:"column:checkUrl" json:"checkUrl"`             // 核销二维码地址
	CheckMessage   string    `gorm:"column:checkMessage" json:"checkMessage"`     // 核销短信
	SynFlag        []uint8   `gorm:"column:synFlag" json:"synFlag"`
	CreateTime     time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime     time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted      []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *ChannelVerificationInfo) TableName() string {
	return "channel_verification_info"
}

// ChannelVerificationInfoColumns get sql column name.获取数据库列名
var ChannelVerificationInfoColumns = struct {
	ID             string
	Channel        string
	OrderNumber    string
	OutOrderNumber string
	CheckCode      string
	CheckURL       string
	CheckMessage   string
	SynFlag        string
	CreateTime     string
	UpdateTime     string
	IsDeleted      string
}{
	ID:             "id",
	Channel:        "channel",
	OrderNumber:    "orderNumber",
	OutOrderNumber: "outOrderNumber",
	CheckCode:      "checkCode",
	CheckURL:       "checkUrl",
	CheckMessage:   "checkMessage",
	SynFlag:        "synFlag",
	CreateTime:     "createTime",
	UpdateTime:     "updateTime",
	IsDeleted:      "isDeleted",
}

// JdShow [...]
type JdShow struct {
	Name string `gorm:"column:name" json:"name"`
	ID   string `gorm:"primaryKey;column:id" json:"-"`
}

// TableName get sql table name.获取数据库表名
func (m *JdShow) TableName() string {
	return "jd_show"
}

// JdShowColumns get sql column name.获取数据库列名
var JdShowColumns = struct {
	Name string
	ID   string
}{
	Name: "name",
	ID:   "id",
}

// JuoooSchedular [...]
type JuoooSchedular struct {
	SchedularID       int     `gorm:"primaryKey;column:schedular_id" json:"-"`
	ShowID            int     `gorm:"column:show_id" json:"showId"`
	ShowTime          string  `gorm:"column:show_time" json:"showTime"`
	Zongdai           []uint8 `gorm:"column:zongdai" json:"zongdai"`
	Method            int8    `gorm:"column:method" json:"method"`
	ShowName          string  `gorm:"column:show_name" json:"showName"`
	CateParentID      int8    `gorm:"column:cate_parent_id" json:"cateParentId"`
	CateParentName    string  `gorm:"column:cate_parent_name" json:"cateParentName"`
	CateChildID       int8    `gorm:"column:cate_child_id" json:"cateChildId"`
	CateChildName     string  `gorm:"column:cate_child_name" json:"cateChildName"`
	ShowPic           string  `gorm:"column:show_pic" json:"showPic"`
	RelateShowID      string  `gorm:"column:relate_show_id" json:"relateShowId"`
	CityID            int     `gorm:"column:city_id" json:"cityId"`
	CityName          string  `gorm:"column:city_name" json:"cityName"`
	VenueID           int     `gorm:"column:venue_id" json:"venueId"`
	VenueName         string  `gorm:"column:venue_name" json:"venueName"`
	Intor             string  `gorm:"column:intor" json:"intor"`
	VedioURL          string  `gorm:"column:vedio_url" json:"vedioUrl"`
	Detail            string  `gorm:"column:detail" json:"detail"`
	Keywords          string  `gorm:"column:keywords" json:"keywords"`
	Tips              string  `gorm:"column:tips" json:"tips"`
	Description       string  `gorm:"column:description" json:"description"`
	SellStatus        int8    `gorm:"column:sell_status" json:"sellStatus"`
	HighPrice         int     `gorm:"column:high_price" json:"highPrice"`
	LowPrice          int8    `gorm:"column:low_price" json:"lowPrice"`
	Shipping          string  `gorm:"column:shipping" json:"shipping"`
	IsThroughTicket   []uint8 `gorm:"column:is_through_ticket" json:"isThroughTicket"`
	CustomShowTime    string  `gorm:"column:custom_show_time" json:"customShowTime"`
	ShowTimeStatrt    []uint8 `gorm:"column:show_time_statrt" json:"showTimeStatrt"`
	ShowTimeEnd       []uint8 `gorm:"column:show_time_end" json:"showTimeEnd"`
	SchedularState    []uint8 `gorm:"column:schedular_state" json:"schedularState"`
	SchedularDelState []uint8 `gorm:"column:schedular_del_state" json:"schedularDelState"`
	ShowState         []uint8 `gorm:"column:show_state" json:"showState"`
	ShowDelState      []uint8 `gorm:"column:show_del_state" json:"showDelState"`
	ShelfTime         string  `gorm:"column:shelf_time" json:"shelfTime"`
}

// TableName get sql table name.获取数据库表名
func (m *JuoooSchedular) TableName() string {
	return "juooo_schedular"
}

// JuoooSchedularColumns get sql column name.获取数据库列名
var JuoooSchedularColumns = struct {
	SchedularID       string
	ShowID            string
	ShowTime          string
	Zongdai           string
	Method            string
	ShowName          string
	CateParentID      string
	CateParentName    string
	CateChildID       string
	CateChildName     string
	ShowPic           string
	RelateShowID      string
	CityID            string
	CityName          string
	VenueID           string
	VenueName         string
	Intor             string
	VedioURL          string
	Detail            string
	Keywords          string
	Tips              string
	Description       string
	SellStatus        string
	HighPrice         string
	LowPrice          string
	Shipping          string
	IsThroughTicket   string
	CustomShowTime    string
	ShowTimeStatrt    string
	ShowTimeEnd       string
	SchedularState    string
	SchedularDelState string
	ShowState         string
	ShowDelState      string
	ShelfTime         string
}{
	SchedularID:       "schedular_id",
	ShowID:            "show_id",
	ShowTime:          "show_time",
	Zongdai:           "zongdai",
	Method:            "method",
	ShowName:          "show_name",
	CateParentID:      "cate_parent_id",
	CateParentName:    "cate_parent_name",
	CateChildID:       "cate_child_id",
	CateChildName:     "cate_child_name",
	ShowPic:           "show_pic",
	RelateShowID:      "relate_show_id",
	CityID:            "city_id",
	CityName:          "city_name",
	VenueID:           "venue_id",
	VenueName:         "venue_name",
	Intor:             "intor",
	VedioURL:          "vedio_url",
	Detail:            "detail",
	Keywords:          "keywords",
	Tips:              "tips",
	Description:       "description",
	SellStatus:        "sell_status",
	HighPrice:         "high_price",
	LowPrice:          "low_price",
	Shipping:          "shipping",
	IsThroughTicket:   "is_through_ticket",
	CustomShowTime:    "custom_show_time",
	ShowTimeStatrt:    "show_time_statrt",
	ShowTimeEnd:       "show_time_end",
	SchedularState:    "schedular_state",
	SchedularDelState: "schedular_del_state",
	ShowState:         "show_state",
	ShowDelState:      "show_del_state",
	ShelfTime:         "shelf_time",
}

// JuoooTicket [...]
type JuoooTicket struct {
	ID           int     `gorm:"primaryKey;column:id" json:"-"`
	Price        int     `gorm:"column:price" json:"price"`
	Stock        int8    `gorm:"column:stock" json:"stock"`
	Name         string  `gorm:"column:name" json:"name"`
	State        int8    `gorm:"column:state" json:"state"`
	Ispackage    []uint8 `gorm:"column:ispackage" json:"ispackage"`
	PackageNum   int8    `gorm:"column:package_num" json:"packageNum"`
	PackagePrice int     `gorm:"column:package_price" json:"packagePrice"`
	SchedularID  int     `gorm:"column:schedular_id" json:"schedularId"`
}

// TableName get sql table name.获取数据库表名
func (m *JuoooTicket) TableName() string {
	return "juooo_ticket"
}

// JuoooTicketColumns get sql column name.获取数据库列名
var JuoooTicketColumns = struct {
	ID           string
	Price        string
	Stock        string
	Name         string
	State        string
	Ispackage    string
	PackageNum   string
	PackagePrice string
	SchedularID  string
}{
	ID:           "id",
	Price:        "price",
	Stock:        "stock",
	Name:         "name",
	State:        "state",
	Ispackage:    "ispackage",
	PackageNum:   "package_num",
	PackagePrice: "package_price",
	SchedularID:  "schedular_id",
}

// ProductChannelSaleRule 商品销售渠道规则表
type ProductChannelSaleRule struct {
	ID           string    `gorm:"primaryKey;column:id" json:"-"`
	ChannelID    string    `gorm:"column:channel_id" json:"channelId"`       // 渠道id
	ItemType     string    `gorm:"column:item_type" json:"itemType"`         // 商品类型（SHOW、ITEM）
	ItemID       string    `gorm:"column:item_id" json:"itemId"`             // 商品类型对应id值
	RuleType     string    `gorm:"column:rule_type" json:"ruleType"`         // 规则类型
	OperatorID   string    `gorm:"column:operator_id" json:"operatorId"`     // 操作人ID
	OperatorName string    `gorm:"column:operator_name" json:"operatorName"` // 操作人名称
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`     // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`     // 更新时间
	IsDeleted    bool      `gorm:"column:is_deleted" json:"isDeleted"`       // 是否删除
}

// TableName get sql table name.获取数据库表名
func (m *ProductChannelSaleRule) TableName() string {
	return "product_channel_sale_rule"
}

// ProductChannelSaleRuleColumns get sql column name.获取数据库列名
var ProductChannelSaleRuleColumns = struct {
	ID           string
	ChannelID    string
	ItemType     string
	ItemID       string
	RuleType     string
	OperatorID   string
	OperatorName string
	CreateTime   string
	UpdateTime   string
	IsDeleted    string
}{
	ID:           "id",
	ChannelID:    "channel_id",
	ItemType:     "item_type",
	ItemID:       "item_id",
	RuleType:     "rule_type",
	OperatorID:   "operator_id",
	OperatorName: "operator_name",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
	IsDeleted:    "is_deleted",
}

// ProductChannelSaleRuleDetail 商品销售渠道规则明细表
type ProductChannelSaleRuleDetail struct {
	ID                string    `gorm:"primaryKey;column:id" json:"-"`
	ChannelSaleRuleID string    `gorm:"column:channel_sale_rule_id" json:"channelSaleRuleId"` // 渠道规则id
	ObjectType        string    `gorm:"column:object_type" json:"objectType"`                 // 对象类型
	ObjectID          string    `gorm:"column:object_id" json:"objectId"`                     // 对象类型对应id值
	SalePrice         float64   `gorm:"column:sale_price" json:"salePrice"`                   // 售价
	Discount          int       `gorm:"column:discount" json:"discount"`                      // 折扣
	TotalStocks       int       `gorm:"column:total_stocks" json:"totalStocks"`               // 库存数量
	Status            string    `gorm:"column:status" json:"status"`                          // 状态
	IsDeleted         bool      `gorm:"column:is_deleted" json:"isDeleted"`                   // 是否删除
	CreateTime        time.Time `gorm:"column:create_time" json:"createTime"`                 // 创建时间
	UpdateTime        time.Time `gorm:"column:update_time" json:"updateTime"`                 // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ProductChannelSaleRuleDetail) TableName() string {
	return "product_channel_sale_rule_detail"
}

// ProductChannelSaleRuleDetailColumns get sql column name.获取数据库列名
var ProductChannelSaleRuleDetailColumns = struct {
	ID                string
	ChannelSaleRuleID string
	ObjectType        string
	ObjectID          string
	SalePrice         string
	Discount          string
	TotalStocks       string
	Status            string
	IsDeleted         string
	CreateTime        string
	UpdateTime        string
}{
	ID:                "id",
	ChannelSaleRuleID: "channel_sale_rule_id",
	ObjectType:        "object_type",
	ObjectID:          "object_id",
	SalePrice:         "sale_price",
	Discount:          "discount",
	TotalStocks:       "total_stocks",
	Status:            "status",
	IsDeleted:         "is_deleted",
	CreateTime:        "create_time",
	UpdateTime:        "update_time",
}

// PushedShow [...]
type PushedShow struct {
	ID       string `gorm:"primaryKey;column:id" json:"-"`
	PushTime string `gorm:"column:push_time" json:"pushTime"`
}

// TableName get sql table name.获取数据库表名
func (m *PushedShow) TableName() string {
	return "pushed_show"
}

// PushedShowColumns get sql column name.获取数据库列名
var PushedShowColumns = struct {
	ID       string
	PushTime string
}{
	ID:       "id",
	PushTime: "push_time",
}

// Supplier 供应商渠道
type Supplier struct {
	ID               string    `gorm:"primaryKey;column:id" json:"-"`                     // 供应商渠道id
	BizCode          string    `gorm:"column:biz_code" json:"bizCode"`                    // 业务编码
	SubBizCode       string    `gorm:"column:sub_biz_code" json:"subBizCode"`             // 二级业务编码
	ChannelName      string    `gorm:"column:channel_name" json:"channelName"`            // 渠道名称
	ChannelShortName string    `gorm:"column:channel_short_name" json:"channelShortName"` // 渠道简称
	ChannelCode      string    `gorm:"column:channel_code" json:"channelCode"`            // 渠道号
	AgentType        string    `gorm:"column:agent_type" json:"agentType"`                // 代理类型
	Contact          string    `gorm:"column:contact" json:"contact"`                     // 联系人
	ContactCellphone string    `gorm:"column:contact_cellphone" json:"contactCellphone"`  // 联系电话
	ContactEmail     string    `gorm:"column:contact_email" json:"contactEmail"`          // 邮箱
	SupplierType     string    `gorm:"column:supplier_type" json:"supplierType"`          // 类型
	EffectiveDate    time.Time `gorm:"column:effective_date" json:"effectiveDate"`        // 有效开始时间
	ExpireDate       time.Time `gorm:"column:expire_date" json:"expireDate"`              // 有效结束时间
	Status           string    `gorm:"column:status" json:"status"`                       // 开启和关闭
	IsDeleted        []uint8   `gorm:"column:is_deleted" json:"isDeleted"`                // 是否删除
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`              // 创建时间
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`              // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *Supplier) TableName() string {
	return "supplier"
}

// SupplierColumns get sql column name.获取数据库列名
var SupplierColumns = struct {
	ID               string
	BizCode          string
	SubBizCode       string
	ChannelName      string
	ChannelShortName string
	ChannelCode      string
	AgentType        string
	Contact          string
	ContactCellphone string
	ContactEmail     string
	SupplierType     string
	EffectiveDate    string
	ExpireDate       string
	Status           string
	IsDeleted        string
	CreateTime       string
	UpdateTime       string
}{
	ID:               "id",
	BizCode:          "biz_code",
	SubBizCode:       "sub_biz_code",
	ChannelName:      "channel_name",
	ChannelShortName: "channel_short_name",
	ChannelCode:      "channel_code",
	AgentType:        "agent_type",
	Contact:          "contact",
	ContactCellphone: "contact_cellphone",
	ContactEmail:     "contact_email",
	SupplierType:     "supplier_type",
	EffectiveDate:    "effective_date",
	ExpireDate:       "expire_date",
	Status:           "status",
	IsDeleted:        "is_deleted",
	CreateTime:       "create_time",
	UpdateTime:       "update_time",
}

// SupplierMappingHistory [...]
type SupplierMappingHistory struct {
	ID            int       `gorm:"primaryKey;column:id" json:"-"`
	ResourceTable string    `gorm:"column:resourceTable" json:"resourceTable"` // 来源映射表
	MappingID     string    `gorm:"column:mappingId" json:"mappingId"`         // 映射表Id
	MappingType   string    `gorm:"column:mappingType" json:"mappingType"`     // 映射类型 三类 演出，场次，票面
	SupplierID    string    `gorm:"column:supplierId" json:"supplierId"`       // 供应商
	SystemID      string    `gorm:"column:systemId" json:"systemId"`           // 平台id
	CreateTime    time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime    time.Time `gorm:"column:updateTime" json:"updateTime"`
	IsDeleted     []uint8   `gorm:"column:isDeleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *SupplierMappingHistory) TableName() string {
	return "supplier_mapping_history"
}

// SupplierMappingHistoryColumns get sql column name.获取数据库列名
var SupplierMappingHistoryColumns = struct {
	ID            string
	ResourceTable string
	MappingID     string
	MappingType   string
	SupplierID    string
	SystemID      string
	CreateTime    string
	UpdateTime    string
	IsDeleted     string
}{
	ID:            "id",
	ResourceTable: "resourceTable",
	MappingID:     "mappingId",
	MappingType:   "mappingType",
	SupplierID:    "supplierId",
	SystemID:      "systemId",
	CreateTime:    "createTime",
	UpdateTime:    "updateTime",
	IsDeleted:     "isDeleted",
}

// SupplierPricingRule 供应商渠道定价规则
type SupplierPricingRule struct {
	ID         string    `gorm:"primaryKey;column:id" json:"-"`        // id
	SupplierID string    `gorm:"column:supplier_id" json:"supplierId"` // 供应商渠道id
	SaleRule   string    `gorm:"column:sale_rule" json:"saleRule"`     // 售价规则
	SaleRatio  float64   `gorm:"column:sale_ratio" json:"saleRatio"`   // 售价系数
	IsDeleted  []uint8   `gorm:"column:is_deleted" json:"isDeleted"`   // 是否删除
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *SupplierPricingRule) TableName() string {
	return "supplier_pricing_rule"
}

// SupplierPricingRuleColumns get sql column name.获取数据库列名
var SupplierPricingRuleColumns = struct {
	ID         string
	SupplierID string
	SaleRule   string
	SaleRatio  string
	IsDeleted  string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	SupplierID: "supplier_id",
	SaleRule:   "sale_rule",
	SaleRatio:  "sale_ratio",
	IsDeleted:  "is_deleted",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// SupplierRegionSeat [...]
type SupplierRegionSeat struct {
	ID                    string    `gorm:"primaryKey;column:id" json:"-"`
	ChannelCode           string    `gorm:"column:channel_code" json:"channelCode"`                       // 渠道编码
	ChannelVenueMappingID string    `gorm:"column:channel_venue_mapping_id" json:"channelVenueMappingId"` // 场馆映射id
	SupplierRegionID      string    `gorm:"column:supplier_region_id" json:"supplierRegionId"`            // 供应商区域id
	FloorID               string    `gorm:"column:floor_id" json:"floorId"`                               // 楼层id
	GraphCol              int       `gorm:"column:graph_col" json:"graphCol"`                             // 物理列
	GraphRow              int       `gorm:"column:graph_row" json:"graphRow"`                             // 物理行
	Col                   string    `gorm:"column:col" json:"col"`                                        // 逻辑列
	Row                   string    `gorm:"column:row" json:"row"`                                        // 逻辑行
	SeatKey               string    `gorm:"column:seat_key" json:"seatKey"`                               // 座位编码
	SeatNo                string    `gorm:"column:seat_no" json:"seatNo"`                                 // 物理座位号
	Remark                string    `gorm:"column:remark" json:"remark"`
	CreateTime            time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime            time.Time `gorm:"column:update_time" json:"updateTime"`
	IsDeleted             []uint8   `gorm:"column:is_deleted" json:"isDeleted"`
}

// TableName get sql table name.获取数据库表名
func (m *SupplierRegionSeat) TableName() string {
	return "supplier_region_seat"
}

// SupplierRegionSeatColumns get sql column name.获取数据库列名
var SupplierRegionSeatColumns = struct {
	ID                    string
	ChannelCode           string
	ChannelVenueMappingID string
	SupplierRegionID      string
	FloorID               string
	GraphCol              string
	GraphRow              string
	Col                   string
	Row                   string
	SeatKey               string
	SeatNo                string
	Remark                string
	CreateTime            string
	UpdateTime            string
	IsDeleted             string
}{
	ID:                    "id",
	ChannelCode:           "channel_code",
	ChannelVenueMappingID: "channel_venue_mapping_id",
	SupplierRegionID:      "supplier_region_id",
	FloorID:               "floor_id",
	GraphCol:              "graph_col",
	GraphRow:              "graph_row",
	Col:                   "col",
	Row:                   "row",
	SeatKey:               "seat_key",
	SeatNo:                "seat_no",
	Remark:                "remark",
	CreateTime:            "create_time",
	UpdateTime:            "update_time",
	IsDeleted:             "is_deleted",
}

// SupplierSettleRule 供应商渠道结算规则
type SupplierSettleRule struct {
	ID              string    `gorm:"primaryKey;column:id" json:"-"`                    // id
	SupplierID      string    `gorm:"column:supplier_id" json:"supplierId"`             // 供应商渠道id
	SettleType      string    `gorm:"column:settle_type" json:"settleType"`             // 结算方式
	SettleTimeType  string    `gorm:"column:settle_time_type" json:"settleTimeType"`    // 结算规则
	SettleInterval  int       `gorm:"column:settle_interval" json:"settleInterval"`     // 结算周期(单位天)
	AgentFeeType    string    `gorm:"column:agent_fee_type" json:"agentFeeType"`        // 代理服务类型
	AgentFeeGmvRate float64   `gorm:"column:agent_fee_gmv_rate" json:"agentFeeGmvRate"` // 整体销售额方式下的代理服务费
	IsDeleted       []uint8   `gorm:"column:is_deleted" json:"isDeleted"`               // 是否删除
	CreateTime      time.Time `gorm:"column:create_time" json:"createTime"`             // 创建时间
	UpdateTime      time.Time `gorm:"column:update_time" json:"updateTime"`             // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *SupplierSettleRule) TableName() string {
	return "supplier_settle_rule"
}

// SupplierSettleRuleColumns get sql column name.获取数据库列名
var SupplierSettleRuleColumns = struct {
	ID              string
	SupplierID      string
	SettleType      string
	SettleTimeType  string
	SettleInterval  string
	AgentFeeType    string
	AgentFeeGmvRate string
	IsDeleted       string
	CreateTime      string
	UpdateTime      string
}{
	ID:              "id",
	SupplierID:      "supplier_id",
	SettleType:      "settle_type",
	SettleTimeType:  "settle_time_type",
	SettleInterval:  "settle_interval",
	AgentFeeType:    "agent_fee_type",
	AgentFeeGmvRate: "agent_fee_gmv_rate",
	IsDeleted:       "is_deleted",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
}

// SupplierShowSettleRule 供应商渠道演出级别结算规则
type SupplierShowSettleRule struct {
	ID         string    `gorm:"primaryKey;column:id" json:"-"`        // id
	SupplierID string    `gorm:"column:supplier_id" json:"supplierId"` // 供应商渠道id
	ShowType   string    `gorm:"column:show_type" json:"showType"`     // 演出类目
	ShowID     string    `gorm:"column:show_id" json:"showId"`         // 演出ID
	ShowName   string    `gorm:"column:show_name" json:"showName"`     // 演出名称
	AgentFee   float64   `gorm:"column:agent_fee" json:"agentFee"`     // 代理服务费率
	IsDeleted  []uint8   `gorm:"column:is_deleted" json:"isDeleted"`   // 是否删除
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *SupplierShowSettleRule) TableName() string {
	return "supplier_show_settle_rule"
}

// SupplierShowSettleRuleColumns get sql column name.获取数据库列名
var SupplierShowSettleRuleColumns = struct {
	ID         string
	SupplierID string
	ShowType   string
	ShowID     string
	ShowName   string
	AgentFee   string
	IsDeleted  string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	SupplierID: "supplier_id",
	ShowType:   "show_type",
	ShowID:     "show_id",
	ShowName:   "show_name",
	AgentFee:   "agent_fee",
	IsDeleted:  "is_deleted",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// TerminalBackup [...]
type TerminalBackup struct {
	ID      int64  `gorm:"primaryKey;column:id" json:"-"`
	Account string `gorm:"column:account" json:"account"`
	Name    string `gorm:"column:name" json:"name"`
	Pwd     string `gorm:"column:pwd" json:"pwd"`
}

// TableName get sql table name.获取数据库表名
func (m *TerminalBackup) TableName() string {
	return "terminal_backup"
}

// TerminalBackupColumns get sql column name.获取数据库列名
var TerminalBackupColumns = struct {
	ID      string
	Account string
	Name    string
	Pwd     string
}{
	ID:      "id",
	Account: "account",
	Name:    "name",
	Pwd:     "pwd",
}

// Test [...]
type Test struct {
	ID      int64  `gorm:"primaryKey;column:id" json:"-"`
	Account string `gorm:"column:account" json:"account"`
	Name    string `gorm:"column:name" json:"name"`
	Pwd     string `gorm:"column:pwd" json:"pwd"`
}

// TableName get sql table name.获取数据库表名
func (m *Test) TableName() string {
	return "test"
}

// TestColumns get sql column name.获取数据库列名
var TestColumns = struct {
	ID      string
	Account string
	Name    string
	Pwd     string
}{
	ID:      "id",
	Account: "account",
	Name:    "name",
	Pwd:     "pwd",
}

// TmpLocation [...]
type TmpLocation struct {
	JdName   string `gorm:"primaryKey;column:jd_name" json:"-"`
	JdName1  string `gorm:"column:jd_name1" json:"jdName1"`
	CityName string `gorm:"column:city_name" json:"cityName"`
	CityID   string `gorm:"column:city_id" json:"cityId"`
}

// TableName get sql table name.获取数据库表名
func (m *TmpLocation) TableName() string {
	return "tmp_location"
}

// TmpLocationColumns get sql column name.获取数据库列名
var TmpLocationColumns = struct {
	JdName   string
	JdName1  string
	CityName string
	CityID   string
}{
	JdName:   "jd_name",
	JdName1:  "jd_name1",
	CityName: "city_name",
	CityID:   "city_id",
}
