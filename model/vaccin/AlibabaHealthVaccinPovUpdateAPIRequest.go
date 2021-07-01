package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinPovUpdateAPIRequest
新增/变更接种点信息 API请求
alibaba.health.vaccin.pov.update

ISV 将疫苗的接种点信息同步到免疫规划中心，提醒用户接种时可提供接种点详情。 */
type AlibabaHealthVaccinPovUpdateAPIRequest struct {
	model.Params
	// 接种点联系电话
	_telephone string
	// 接种点具体地址
	_address string
	// 接种点介绍
	_description string
	// 接种点编码
	_povNo string
	// 接种点名称
	_povName string
	// 服务时间
	_businessTime string
}

// New
