package config

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("Given an environment with no environment variables set", t, func() {

		Convey("When the config is retrieved", func() {
			cfg, err := Get()
			Convey("Then there should be no error returned", func() {
				So(err, ShouldBeNil)
			})

			Convey("The values should be set to the expected defaults", func() {
				So(cfg.ConsumerQueue, ShouldEqual, "")
				So(cfg.ConsumerQueueURL, ShouldEqual, "")
				So(cfg.ProducerQueue, ShouldEqual, "")
				So(cfg.QueueRegion, ShouldEqual, "eu-west-1")
				So(cfg.VerificationKey, ShouldEqual, "")
				So(cfg.DeploymentRoot, ShouldEqual, "")
				So(cfg.NomadEndpoint, ShouldEqual, "http://localhost:4646")
				So(cfg.NomadToken, ShouldEqual, "")
				So(cfg.NomadCACert, ShouldEqual, "")
				So(cfg.NomadTLSSkipVerify, ShouldBeFalse)
				So(cfg.S3DeploymentRegion, ShouldEqual, "eu-west-1")
				So(cfg.DeploymentTimeout, ShouldEqual, time.Second*60*20)
				So(cfg.BindAddr, ShouldEqual, ":24300")
				So(cfg.HealthcheckInterval, ShouldEqual, time.Second*30)
				So(cfg.HealthcheckCriticalTimeout, ShouldEqual, time.Second*10)
				So(cfg.PrivateKey, ShouldEqual, "")
				So(cfg.S3SecretsRegion, ShouldEqual, "eu-west-1")
			})
		})
	})
}