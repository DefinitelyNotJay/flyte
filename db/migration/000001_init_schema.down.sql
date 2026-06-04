-- ลบตารางที่มีความสัมพันธ์ (Foreign Keys) กับตารางอื่นก่อน
DROP TABLE IF EXISTS "notifications";
DROP TABLE IF EXISTS "post_hashtags";
DROP TABLE IF EXISTS "media";
DROP TABLE IF EXISTS "likes";
DROP TABLE IF EXISTS "posts";
DROP TABLE IF EXISTS "follows";
DROP TABLE IF EXISTS "refresh_tokens";

-- ลบตารางหลักที่ไม่มี Foreign Key ไปตารางอื่น (แต่อาจถูกตารางอื่นอ้างอิง)
DROP TABLE IF EXISTS "hashtags";
DROP TABLE IF EXISTS "users";

-- ลบ Custom Types (ENUM) เป็นลำดับสุดท้าย
DROP TYPE IF EXISTS "notification_type";