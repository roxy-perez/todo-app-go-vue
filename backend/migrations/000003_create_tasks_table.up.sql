CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    sprint_id INT REFERENCES sprints(id) ON DELETE SET NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) DEFAULT 'pending', -- pending, in_progress, done
    priority VARCHAR(50) DEFAULT 'medium', -- low, medium, high
    tags TEXT[], -- array de etiquetas
    due_date DATE,
    created_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT now()
);