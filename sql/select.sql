-- name: select-purchase
SELECT vendor, articleId, bulk FROM purchase;

-- name: select-purchase-by-articleId
SELECT vendor, articleId, bulk FROM purchase WHERE articleId=?;

-- name: search-vendor
SELECT vendor, levenshteinDistance(vendor,?) as ld FROM purchase WHERE ld <= 10 ORDER BY ld;

-- name: search-purchase-by-time
SELECT vendor, articleId, bulk FROM purchase WHERE created >= ? AND created <= ?;