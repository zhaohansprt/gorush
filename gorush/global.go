package gorush

import (
	"crypto/tls"

	"github.com/zhaohansprt/gorush/config"
	"github.com/appleboy/gorush/storage"

	"github.com/appleboy/go-fcm"
	apns "github.com/sideshow/apns2"
	"github.com/sirupsen/logrus"
)

var (
	// PushConf is gorush config
	PushConf config.ConfYaml
	// QueueNotification is chan type
	QueueNotification chan PushNotification
	// CertificatePemIos is ios certificate file
	CertificatePemIos tls.Certificate
	//multi-cert apns client
	ApnsClientMap map[int] *apns.Client
	// ApnsClient is apns client
	ApnsClient *apns.Client
	// FCMClient is apns client
	FCMClient *fcm.Client
	// LogAccess is log server request log
	LogAccess *logrus.Logger
	// LogError is log server error log
	LogError *logrus.Logger
	// StatStorage implements the storage interface
	StatStorage storage.Storage
)
