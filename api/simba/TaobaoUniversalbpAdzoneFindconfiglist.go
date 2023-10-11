package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpAdzoneFindconfiglist 查询所有可用资源包信息
// taobao.universalbp.adzone.findconfiglist
//
// 查询该场景下，所有可用的资源包，可能存在数据重复。但是针对不同子场景和推广设置，可以选用的资源包有差异，建议关注补充文档，或者根据bp前端的限制，进行传参。
func TaobaoUniversalbpAdzoneFindconfiglist(clt *core.SDKClient, req *simba.TaobaoUniversalbpAdzoneFindconfiglistAPIRequest, session string) (*simba.TaobaoUniversalbpAdzoneFindconfiglistAPIResponse, error) {
	var resp simba.TaobaoUniversalbpAdzoneFindconfiglistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
