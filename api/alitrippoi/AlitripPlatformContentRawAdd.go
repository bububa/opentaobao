package alitrippoi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitrippoi"
)

// Alitripplatformcontentrawadd 穷游内容写入接口
// alitrip.platform.content.raw.add
//
// 穷游内容写入飞猪接口
func Alitripplatformcontentrawadd(clt *core.SDKClient, req *alitrippoi.AlitripplatformcontentrawaddAPIRequest, session string) (*alitrippoi.AlitripplatformcontentrawaddAPIResponse, error) {
	var resp alitrippoi.AlitripplatformcontentrawaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
