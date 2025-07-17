const fs = require("fs");
const path = require("path");

const FILE_FOLDER = path.join(__dirname, "..", "file");
const VALID_USERNAME = "admin";
const VALID_PASSWORD = "admin";

if (!fs.existsSync(FILE_FOLDER)) fs.mkdirSync(FILE_FOLDER);

module.exports = {
  FILE_FOLDER,
  VALID_USERNAME,
  VALID_PASSWORD,
};
