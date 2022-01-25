CREATE TABLE IF NOT EXISTS orgs (
    id SERIAL NOT NULL,
    name TEXT NOT NULL,
    CONSTRAINT pk_orgs PRIMARY KEY(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL NOT NULL,
    org_id INT NOT NULL,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    username VARCHAR(150) NOT NULL UNIQUE,
    PASSWORD VARCHAR(256) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    picture VARCHAR(256) DEFAULT 'https://place.dog/300/300',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT pk_users PRIMARY KEY(id),
    CONSTRAINT fk_users_org FOREIGN KEY(org_id) REFERENCES orgs(id)
);

CREATE TABLE IF NOT EXISTS points (
    id SERIAL NOT NULL,
    user_id INT NOT NULL,
    points numeric(12, 2) NOT NULL DEFAULT 0.0,
    metadata jsonb NULL,
    assigned_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT pk_points PRIMARY KEY(id),
    CONSTRAINT fk_user_points FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS badges (
    id SERIAL NOT NULL,
    org_id INT NOT NULL,
    name TEXT NOT NULL,
    image TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT pk_badges PRIMARY KEY(id),
    CONSTRAINT fk_badges_org FOREIGN KEY(org_id) REFERENCES orgs(id)
);

CREATE TABLE IF NOT EXISTS user_badges (
    user_id INT NOT NULL,
    badge_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT pk_user_badges PRIMARY KEY(user_id, badge_id),
    CONSTRAINT fk_user_user_badges FOREIGN KEY(badge_id) REFERENCES badges(id),
    CONSTRAINT fk_badge_user_badges FOREIGN KEY(user_id) REFERENCES users(id)
);