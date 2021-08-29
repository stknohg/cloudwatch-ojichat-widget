# cloudwatch-ojichat-widget

[ojichat](https://github.com/greymd/ojichat)を使ったシンプルなCloudWatch custom widgetです。  
[CloudWatch custom widget](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/add_custom_widget_dashboard.html)のサンプルとしてお使いください。 

## 使い方

Lambda関数のデプロイで楽をしかったため[AWS Serverless Application Model(SAM)](https://aws.amazon.com/jp/serverless/sam/)を使っています。  
SAMおよびGo言語の開発環境を整えたうえで`sam build`および`sam deploy`を使ってデプロイしてください。  

```bash
# samconfig.toml に必要な設定を記載済の前提
sam build && sam deploy
```

### デモ用ダッシュボード

CloudFormation Templateの`DoCreateExampleDashboard`パラメーターを`Yes`にするとデモ用ダッシュボードを作成します。  
とりあえず試したい場合はこのパラメーターを設定すると良いでしょう。  

```bash
# samconfig.toml に必要な設定を記載済の前提
sam build && sam deploy --parameter-overrides DoCreateDemoDashboard="Yes"
```

## Thanks

本ソフトウェアは以下を参考にしています。

* [greymd/ojichat](https://github.com/greymd/ojichat)
* [mokocm/serverless-ojichat](https://github.com/mokocm/serverless-ojichat)

## License

* [MIT](./LICENSE)
