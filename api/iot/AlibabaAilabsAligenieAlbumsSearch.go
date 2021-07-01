package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

/* AlibabaAilabsAligenieAlbumsSearch
查询专辑
alibaba.ailabs.aligenie.albums.search

搜索类目下的专辑信息 */
func AlibabaAilabsAligenieAlbumsSearch(clt *core.SDKClient, req *iot.AlibabaAilabsAligenieAlbumsSearchAPIRequest, session string) (*iot.AlibabaAilabsAligenieAlbumsSearchAPIResponse, error) {
	var resp iot.AlibabaAilabsAligenieAlbumsSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
