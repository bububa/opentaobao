package dengta

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dengta"
)

// Alibabapicturesdengtaordereffectimport 天下秀订单数据导入
// alibaba.pictures.dengta.order.effect.import
//
// 提供接口给天下秀，天下秀订单数据效果生成时回流到灯塔系统
func Alibabapicturesdengtaordereffectimport(clt *core.SDKClient, req *dengta.AlibabapicturesdengtaordereffectimportAPIRequest, session string) (*dengta.AlibabapicturesdengtaordereffectimportAPIResponse, error) {
	var resp dengta.AlibabapicturesdengtaordereffectimportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
