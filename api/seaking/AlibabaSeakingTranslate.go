package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

/* AlibabaSeakingTranslate
MT定制接口
alibaba.seaking.translate

MT定制接口 */
func AlibabaSeakingTranslate(clt *core.SDKClient, req *seaking.AlibabaSeakingTranslateAPIRequest, session string) (*seaking.AlibabaSeakingTranslateAPIResponse, error) {
	var resp seaking.AlibabaSeakingTranslateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
