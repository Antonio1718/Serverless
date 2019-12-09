'use strict';

const serverless = require('serverless-http');
// const express = require('express')
// const app = express()

const AWS = require('aws-sdk');
const sqs = new AWS.SQS({region: 'eu-west-3'});
const fetch = require('node-fetch');

const QUEUE_URL = `https://sqs.eu-west-3.amazonaws.com/300334162006/asecond`;
// Get Message in XIPE API and sending new SQS File d'attente  
exports.handler = (event, context, callback) => {
  let response;
  console.log('********type*********: ',typeof event);
  console.log('===========: ', event);
  // Send Xipe-Api Validation
  fetch('http://35.181.155.125/')
    .then(res => res.json())
    .then(resp => {
      // Case success in API
      const params = {
        MessageBody: resp.service,
        QueueUrl: QUEUE_URL
      };
      sqs.sendMessage(params, function(err, data) {      
        if (err) {
          console.log('error:', 'Fail Send Message' + err);
           response = {
            statusCode: 500,
              body: JSON.stringify({
              message: err
            })
          };
        } else {
          console.log('data:', data);
          response = {
            statusCode: 200,
            body: JSON.stringify({
              message: data,
              body:resp
            })
          };
        }
        callback(null, response)  
      })
    })   
};

