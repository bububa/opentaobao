package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

/* TaobaoTanxQualificationAdvertiserAdd
新增广告主接口
taobao.tanx.qualification.advertiser.add

外部dsp调用接口时会根据广告主名称和类型在tanx系统中新增一个广告主 */
func TaobaoTanxQualificationAdvertiserAdd(clt *core.SDKClient, req *tanx.TaobaoTanxQualificationAdvertiserAddAPIRequest, session string) (*tanx.TaobaoTanxQualificationAdvertiserAddAPIResponse, error) {
	var resp tanx.TaobaoTanxQualificationAdvertiserAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
