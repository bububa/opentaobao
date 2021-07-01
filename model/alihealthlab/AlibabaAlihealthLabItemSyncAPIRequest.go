package alihealthlab

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthLabItemSyncAPIRequest
阿里健康检验检测商品发布 API请求
alibaba.alihealth.lab.item.sync

iSV发布检验检测商品基本信息给健康，内部关联一个淘宝商品或SKU */
type AlibabaAlihealthLabItemSyncAPIRequest struct {
	model.Params
	// 检验检测项目isv侧编码
	_isvItemCode string
	// 检验检测项目名称
	_name string
	// 采购价，单位分
	_costPrice int64
	// 性别限制 MALE  FEMALE
	_genderRestriction string
	// 婚否状态限制 MARRIED UNMARRIED
	_maritalStatusRestriction string
	// 额外的属性
	_extraAttributes string
	// 报告产出预计需要的时长，单位毫秒
	_reportGenerationTimeDuration int64
	// 打包的子项目
	_subItems []LabSubItemBrief
	// EFFECTIVE 项目有效， INVALID 项目无效
	_isvItemStatus string
	// 项目关联的门店
	_relatedIsvStoreCodes []string
	// 项目介绍
	_intro string
}

// New
