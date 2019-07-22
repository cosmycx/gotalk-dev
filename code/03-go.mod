module msg-entry-serv

require (
	github.com/Shopify/sarama v1.23.0
	github.com/go-ini/ini v1.44.0 // indirect
	github.com/minio/minio-go v6.0.14+incompatible
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	user-serv v0.0.0 // Local
)

replace user-serv => ../user-serv