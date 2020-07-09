package functions

import (
	"encoding/json"
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

	// var report2 util.Report

	var report util.Report
	json.Unmarshal(b, &report)

	// err := json.NewDecoder(strings.NewReader(string(b))).Decode(&report)
	// if err != nil {
	// 	fmt.Println(report)
	// }
	timestamp := time.Now().Unix()
	report.Timestamp = strconv.FormatInt(timestamp, 10)
	pow.StoreReport(report, ctx)

	return nil
}
