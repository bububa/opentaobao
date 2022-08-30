package alihouse

import (
	"github.com/bububa/opentaobao/model"
)

// TradeToolDto 结构体
type TradeToolDto struct {
	// 工具唯一id，如购房登记id
	OuterToolId string `json:"outer_tool_id,omitempty" xml:"outer_tool_id,omitempty"`
	// 项目名称
	ProjectName string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// 链接名称
	LinkName string `json:"link_name,omitempty" xml:"link_name,omitempty"`
	// 链接地址
	LinkUrl string `json:"link_url,omitempty" xml:"link_url,omitempty"`
	// 状态显示名称
	StatusShowName string `json:"status_show_name,omitempty" xml:"status_show_name,omitempty"`
	// 开始时间
	BeginTime string `json:"begin_time,omitempty" xml:"begin_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 外部门店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 是否显示
	IsShow *model.File `json:"is_show,omitempty" xml:"is_show,omitempty"`
	// 工具类型，1. 购房登记
	ToolType *model.File `json:"tool_type,omitempty" xml:"tool_type,omitempty"`
	// 项目状态,0-未开始、1-进行中、2-审核中、3-已结束
	ProjectStatus *model.File `json:"project_status,omitempty" xml:"project_status,omitempty"`
	// 是否为测试数据
	IsTest *model.File `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// ETC请求时间版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}
