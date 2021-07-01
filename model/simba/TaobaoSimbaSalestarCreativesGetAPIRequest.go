package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaSalestarCreativesGetAPIRequest
（新）批量获取创意 API请求
taobao.simba.salestar.creatives.get

取得一个推广组的所有创意或者根据一个创意Id列表取得一组创意；<br/>如果同时提供了推广组Id和创意id列表，则优先使用推广组Id； */
type TaobaoSimbaSalestarCreativesGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 创意Id数组，最多200个
	_creativeIds []int64
	// 推广组Id
	_adgroupId int64
}

// New
