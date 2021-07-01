package wangwang

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wangwang"
)

/* TaobaoWangwangAbstractAddword
增加关键词
taobao.wangwang.abstract.addword

增加关键词，只支持json返回 */
func TaobaoWangwangAbstractAddword(clt *core.SDKClient, req *wangwang.TaobaoWangwangAbstractAddwordAPIRequest, session string) (*wangwang.TaobaoWangwangAbstractAddwordAPIResponse, error) {
	var resp wangwang.TaobaoWangwangAbstractAddwordAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
