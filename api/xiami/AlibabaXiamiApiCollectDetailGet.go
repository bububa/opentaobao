package xiami

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiami"
)

/* AlibabaXiamiApiCollectDetailGet
虾米音乐精选集详情接口
alibaba.xiami.api.collect.detail.get

虾米音乐精选集详情接口 */
func AlibabaXiamiApiCollectDetailGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiCollectDetailGetAPIRequest, session string) (*xiami.AlibabaXiamiApiCollectDetailGetAPIResponse, error) {
	var resp xiami.AlibabaXiamiApiCollectDetailGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
