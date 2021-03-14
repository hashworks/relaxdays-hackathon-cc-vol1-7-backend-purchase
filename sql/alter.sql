-- name: insert-purchase
INSERT INTO purchase(vendor, articleId, bulk, created, priceInCents) VALUES(?, ?, ?, strftime('%s', 'now', 'localtime'), ?);