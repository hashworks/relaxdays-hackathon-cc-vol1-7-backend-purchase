-- name: insert-purchase
INSERT INTO purchase(vendor, articleId, bulk, created) VALUES(?, ?, ?, strftime('%s', 'now', 'localtime'));