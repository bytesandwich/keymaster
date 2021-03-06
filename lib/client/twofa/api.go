// Package twofa contains routines for getting short lived certificate.
package twofa

import (
	"crypto"
	"crypto/x509"
	"flag"
	"time"

	"github.com/Symantec/Dominator/lib/log"
)

var (
	// Duration of generated cert. Default 16 hours.
	Duration = flag.Duration("duration", 16*time.Hour, "Duration of the requested certificates in golang duration format (ex: 30s, 5m, 12h)")
	// If set, Do not use U2F as second factor
	noU2F = flag.Bool("noU2F", false, "Don't use U2F as second factor")
	// If set, Do not use VIPAccess as second factor.
	noVIPAccess = flag.Bool("noVIPAccess", false, "Don't use VIPAccess as second factor")
)

// GetCertFromTargetUrls gets a signed cert from the given target URLs.
func GetCertFromTargetUrls(
	signer crypto.Signer,
	userName string,
	password []byte,
	targetUrls []string,
	rootCAs *x509.CertPool,
	skipu2f bool,
	logger log.DebugLogger) (sshCert []byte, x509Cert []byte, kubernetesCert []byte, err error) {
	return getCertFromTargetUrls(
		signer, userName, password, targetUrls, rootCAs, skipu2f, logger)
}
