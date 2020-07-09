#!/bin/sh
cd functions && gcloud functions deploy submitReport --runtime go113 --trigger-http --entry-point SubmitReportHandler --allow-unauthenticated
