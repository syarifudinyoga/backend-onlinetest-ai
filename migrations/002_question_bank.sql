CREATE TABLE IF NOT EXISTS question_bank (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    question_type VARCHAR(50) NOT NULL,
    -- multiple_choice | essay | true_false | coding | file_upload | audio | image

    question_text TEXT NOT NULL,

    explanation TEXT,

    metadata JSONB DEFAULT '{}',

    created_by UUID,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- OPTIONS (khusus MCQ)
CREATE TABLE IF NOT EXISTS question_options (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    question_id UUID REFERENCES question_bank(id) ON DELETE CASCADE,

    option_text TEXT NOT NULL,

    is_correct BOOLEAN DEFAULT FALSE
);

-- TAGGING (biar AI-ready)
CREATE TABLE IF NOT EXISTS question_tags (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    question_id UUID REFERENCES question_bank(id) ON DELETE CASCADE,

    tag VARCHAR(50) NOT NULL
);