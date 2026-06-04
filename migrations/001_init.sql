-- BUAT MANUAL AJA DI DB
-- CREATE TABLE IF NOT EXISTS schema_migrations (
--     id SERIAL PRIMARY KEY,
--     filename TEXT UNIQUE,
--     executed_at TIMESTAMP DEFAULT NOW()
-- );

-- =========================
-- USERS
-- =========================
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(150) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role VARCHAR(50) NOT NULL, -- admin, participant, reviewer
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_roles (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    role VARCHAR(50) NOT NULL -- admin, participant, reviewer
);

-- =========================
-- AUTH (JWT REFRESH OPTIONAL)
-- =========================
-- CREATE TABLE auth_sessions (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     user_id UUID REFERENCES users(id) ON DELETE CASCADE,
--     refresh_token TEXT,
--     expires_at TIMESTAMP,
--     created_at TIMESTAMP DEFAULT NOW()
-- );

-- =========================
-- QUESTION BANK
-- =========================
-- CREATE TABLE question_bank (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     question_type VARCHAR(50) NOT NULL,
--     question_text TEXT NOT NULL,

--     -- fleksibel untuk AI scoring & future type soal
--     metadata JSONB,

--     created_at TIMESTAMP DEFAULT NOW(),
--     updated_at TIMESTAMP DEFAULT NOW()
-- );

-- CREATE TABLE question_options (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     question_id UUID REFERENCES question_bank(id) ON DELETE CASCADE,
--     option_text TEXT NOT NULL,
--     is_correct BOOLEAN DEFAULT FALSE
-- );

-- =========================
-- TEST PACKAGE
-- =========================
-- CREATE TABLE test_packages (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     name VARCHAR(150) NOT NULL,
--     description TEXT,
--     duration_minutes INT DEFAULT 60,
--     created_at TIMESTAMP DEFAULT NOW()
-- );

-- CREATE TABLE test_package_questions (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     package_id UUID REFERENCES test_packages(id) ON DELETE CASCADE,
--     question_id UUID REFERENCES question_bank(id) ON DELETE CASCADE,
--     weight INT DEFAULT 1
-- );

-- =========================
-- EVENT (RECRUITMENT / TEST SESSION GROUP)
-- =========================
-- CREATE TABLE events (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     name VARCHAR(150) NOT NULL,
--     start_date TIMESTAMP,
--     end_date TIMESTAMP,
--     created_at TIMESTAMP DEFAULT NOW()
-- );

-- CREATE TABLE event_users (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     event_id UUID REFERENCES events(id) ON DELETE CASCADE,
--     user_id UUID REFERENCES users(id) ON DELETE CASCADE
-- );

-- CREATE TABLE event_packages (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     event_id UUID REFERENCES events(id) ON DELETE CASCADE,
--     package_id UUID REFERENCES test_packages(id) ON DELETE CASCADE
-- );

-- =========================
-- TEST SESSION
-- =========================
-- CREATE TABLE test_sessions (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     user_id UUID REFERENCES users(id) ON DELETE CASCADE,
--     package_id UUID REFERENCES test_packages(id) ON DELETE CASCADE,
--     event_id UUID REFERENCES events(id),

--     started_at TIMESTAMP DEFAULT NOW(),
--     finished_at TIMESTAMP,
--     status VARCHAR(50) DEFAULT 'ongoing'
-- );

-- CREATE TABLE test_answers (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     session_id UUID REFERENCES test_sessions(id) ON DELETE CASCADE,
--     question_id UUID REFERENCES question_bank(id),
--     answer TEXT,
--     score NUMERIC DEFAULT 0,

--     -- penting buat AI scoring nanti
--     ai_score JSONB,

--     created_at TIMESTAMP DEFAULT NOW()
-- );

-- -- =========================
-- -- REPORTING (PRECOMPUTED / CACHE)
-- -- =========================
-- CREATE TABLE test_reports (
--     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--     session_id UUID REFERENCES test_sessions(id) ON DELETE CASCADE,
--     total_score NUMERIC DEFAULT 0,
--     detail JSONB,
--     generated_at TIMESTAMP DEFAULT NOW()
-- );