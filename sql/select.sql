-- name: select-purchase
SELECT vendor, articleId, bulk FROM purchase;

-- name: select-purchase-by-articleId
SELECT vendor, articleId, bulk FROM purchase WHERE articleId=?;