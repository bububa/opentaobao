package autonavi

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/autonavi"
)

// Alibabaautonaviapitrafficboardimageget 交通看板-栅格情报获取
// alibaba.autonavi.api.trafficboard.image.get
//
// 获取指定情报板ID的二进制数据（图片）
func Alibabaautonaviapitrafficboardimageget(clt *core.SDKClient, req *autonavi.AlibabaautonaviapitrafficboardimagegetAPIRequest, session string) (*autonavi.AlibabaautonaviapitrafficboardimagegetAPIResponse, error) {
	var resp autonavi.AlibabaautonaviapitrafficboardimagegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
