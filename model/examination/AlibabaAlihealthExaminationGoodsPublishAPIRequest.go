package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationGoodsPublishAPIRequest
体检机构对接_商品发布／更新 API请求
alibaba.alihealth.examination.goods.publish

体检机构对接_商品发布／更新 */
type AlibabaAlihealthExaminationGoodsPublishAPIRequest struct {
	model.Params
	// 商品id，机构保证全局唯一
	_groupId string
	// 商品名称
	_groupName string
	// 套餐列表
	_packageList []Package
	// 操作类型: publish=发布，update=更新
	_type string
	// 最多200个字，界面对应商品详情页描述
	_goodsDesc string
	// 最多256个字，界面对应列表文字
	_targetGroup string
	// 联调中正式上线前标签给B；联调后正式上线后标签给C
	_label string
	// 商品类目，1：体检 ，2：核酸，4 ：健康证
	_categoryId string
	// 0自营商品，1平台商品
	_mode string
	// 类目ID，填入叶子类目ID，儿童体检: 20210204000004, 中青年体检: 20210204000005, 老年体检: 20210204000006, 证件体检（含入职）: 20210204000007, 核酸检测（到店服务）: 20210204000008, 专科服务（不包含核酸检测）: 20210204000009, 上门检测: 202102040000010, 上门护理: 202102040000011, 上门体检 202102040000012
	_backendCategoryId int64
}

// New
