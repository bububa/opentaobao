package metadata

const (
	DOC_CENTER_URL  = "https://open.taobao.com/docCenter"                                                                                         // 文档中心链接，用于获取_tb_token_ cookie
	API_CATELOG_URL = "https://open.taobao.com/handler/document/getApiCatelogConfig.json?scopeId=&treeId=&docId=285&docType=2&tag=&_tb_token_=%s" // 文档类目地址
	DOC_URL         = "https://open.taobao.com/handler/document/getDocument.json?isEn=false&treeId=&docId=%d&docType=2&_tb_token_=%s"             // 单API文档地址
	DOC_LINK        = "https://open.taobao.com/API.htm?docId=%d&docType=2"                                                                        // API文档链接
)
