-- name: GetNote :one
SELECT * FROM notes
WHERE id = $1 limit 1;


-- name: CreateNote :one
INSERT INTO notes (
    createdBy,
    content
) VALUES (
    $1,
    $2
)
RETURNING *;

-- name: UpdateLLMResp :one
UPDATE notes
    set llmResp = $2,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteNote :exec
DELETE FROM notes
WHERE id = $1;

