package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationItemsPublishAPIRequest
单项/加项包信息同步 API请求
alibaba.alihealth.examination.items.publish

体检机构对接_单项/加项包信息发布／更新 */
type AlibabaAlihealthExaminationItemsPublishAPIRequest struct {
	model.Params
	// 商品id，机构保证全局唯一
	_groupId string
	// 套餐列表
	_isvPackages []IsvPackage
	// 单项之间关系
	_isvItemRelationDTOS []IsvItemRelationDto
	// 体检机构标识
	_hospitalCodes []string
	// 加项包列表
	_isvItemPackDTOS []IsvItemPackDto
	// 单项信息列表
	_isvItemDTOS []IsvItemDto
	// 加项包关系列表
	_isvPackRelationDTOS []IsvPackRelationDto
}

// New
