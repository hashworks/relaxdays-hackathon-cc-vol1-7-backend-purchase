-- name: create-table-purchase
CREATE TABLE IF NOT EXISTS purchase(
  purchaseId INTEGER PRIMARY KEY AUTOINCREMENT,
  vendor VARCHAR NOT NULL,
  articleId INTEGER,
  bulk INTEGER
);