const express = require("express");
const bodyParser = require("body-parser");

const app = express();
app.use(bodyParser.json());

app.post("/cashapp/deposit", (req, res) => {
  console.log("Received CashApp deposit:", req.body);
  res.json({ status: "ok", txid: "fake123" });
});

app.listen(8084, () => console.log("CashApp proxy running on :8084"));
