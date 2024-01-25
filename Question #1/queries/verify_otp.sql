-- name: VerifyOTP :one
SELECT id FROM users WHERE phone_number = $1 AND otp = $2 AND otp_expiration_time > NOW();
