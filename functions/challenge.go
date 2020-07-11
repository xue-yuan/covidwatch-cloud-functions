package functions

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"upload-token.functions/internal/pow"
	"upload-token.functions/internal/util"
)

// ChallengeHandler is a handler for the /challenge endpoint.
var ChallengeHandler = util.MakeHTTPHandler(challengeHandler)

// SubmitReportHandler is a handler for the /submitReport endpoint.
var SubmitReportHandler = util.MakeHTTPHandler(submitReportHandler)

func challengeHandler(ctx *util.Context) util.StatusError {
	if err := util.ValidateRequestMethod(ctx, "GET", ""); err != nil {
		return err
	}

	c, err := pow.GenerateChallenge(ctx)
	if err != nil {
		return util.NewInternalServerError(err)
	}
	json.NewEncoder(ctx.HTTPResponseWriter()).Encode(c)

	return nil
}

func submitReportHandler(ctx *util.Context) util.StatusError {
	if err := util.ValidateRequestMethod(ctx, "POST", ""); err != nil {
		return err
	}
	defer ctx.HTTPRequest().Body.Close()
	b, _ := ioutil.ReadAll(ctx.HTTPRequest().Body)

	fmt.Println(b)
	var report util.Report
	json.Unmarshal(b, &report)

	eI := report.EndIndex
	mD, _ := base64.StdEncoding.DecodeString(report.MemoData)
	mT := report.MemoType
	rB, _ := base64.StdEncoding.DecodeString(report.ReportVerificationPublicKeyBytes)
	sB, _ := base64.StdEncoding.DecodeString(report.SignatureBytes)
	sI := report.StartIndex
	tB, _ := base64.StdEncoding.DecodeString(report.TemporaryContactKeyBytes)

	timestamp := time.Now().Unix()
	report.Timestamp = strconv.FormatInt(timestamp, 10)

	mReport := map[string]interface{}{
		"end_index":                            eI,
		"memo_data":                            mD,
		"memo_type":                            mT,
		"report_verification_public_key_bytes": rB,
		"signature_bytes":                      sB,
		"start_index":                          sI,
		"temporary_contact_key_bytes":          tB,
		"timestamp":                            report.Timestamp,
	}

	// fmt.Println(mReport)
	pow.StoreReport(report, ctx, mReport)

	return nil
}
