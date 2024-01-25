-- name: GenerateOTP :one
UPDATE users SET otp = $2, otp_expiration_time = $3 WHERE phone_number = $1 RETURNING *;
