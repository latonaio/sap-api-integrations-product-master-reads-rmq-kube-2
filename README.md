# sap-api-integrations-product-master-reads-rmq-kube

sap-api-integrations-product-master-reads-rmq-kube は、外部システム(特にエッジコンピューティング環境)を SAP と統合することを目的に、SAP API で品目マスタデータを取得するマイクロサービスです。  
sap-api-integrations-product-master-reads-rmq-kube には、サンプルの API Json フォーマットが含まれています。  
sap-api-integrations-product-master-reads-rmq-kube は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版 API を利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_PRODUCT_SRV_0001/overview

## 動作環境

sap-api-integrations-product-master-reads-rmq-kube は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （必須）  
・ AION のリソース （推奨)  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  
・ RabbitMQ  

## クラウド環境での利用

sap-api-integrations-product-master-reads-rmq-kube は、外部システムがクラウド環境である場合に SAP と統合するときにおいても、利用可能なように設計されています。

## Output

本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータが JSON 形式で出力されます。  
以下の sample.json の例は、SAP 品目マスタ の 一般データ が取得された結果の JSON の例です。  
以下の項目のうち、"BaseUnit" ～ "WeightUnit" は、/SAP_API_Caller/formatter.go 内 の type Material struct {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。

```
{
	"BaseUnit": "AU",
	"Division": "00",
	"GrossWeight": "0.000",
	"Material": "A001",
	"ProductGroup": "A001",
	"ProductStandardID": "",
	"Product_desc": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PRODUCT_SRV/A_Product('A001')/to_Description",
	"SizeOrDimensionText": "",
	"ValidityStartDate": "",
	"WeightUnit": "KG",
	"cursor": "/Users/kyotatashiro/go/src/sap-api-integrations-product-master-reads-rmq-kube/SAP_API_Caller.(*SAPAPICaller).Product",
	"level": "INFO",
	"time": "2021-11-26T15:50:30.156715+09:00"
}

```
