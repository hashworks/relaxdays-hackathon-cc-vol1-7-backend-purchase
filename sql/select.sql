-- name: select-purchase
SELECT vendor, articleId, bulk, priceInCents / 100.0 as priceInEuro FROM purchase;

-- name: select-purchase-by-articleId
SELECT vendor, articleId, bulk, priceInCents / 100.0 as priceInEuro FROM purchase WHERE articleId=?;

-- name: search-vendor
SELECT DISTINCT vendor, levenshteinDistance(vendor,?) as ld FROM purchase WHERE ld <= 10 ORDER BY ld;

-- name: search-purchase-by-time
SELECT vendor, articleId, bulk, priceInCents / 100.0 as priceInEuro FROM purchase WHERE created >= ? AND created <= ?;

-- name: search-price-over-time
SELECT priceInCents / 100.0 as priceInEuro, created FROM purchase WHERE articleId = ? ORDER BY created;