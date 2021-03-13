-- name: select-purchase
SELECT vendor, articleId, bulk FROM purchase;

-- name: select-purchase-by-articleId
SELECT vendor, articleId, bulk FROM purchase WHERE articleId=?;

-- name: search-purchase-by-vendor
SELECT vendor, articleId, bulk, levenshteinDistance(vendor,?) as ld FROM purchase WHERE ld <= 10 ORDER BY ld;