-- PostgreSQL solution
SELECT
    ROUND(COUNT(logged_day_after)::decimal / COUNT(player_id), 2) AS fraction
FROM (
    SELECT
        a.player_id,
        MAX(
            CASE
                WHEN a.event_date - t.min_event_date = 1
                THEN 1
                ELSE NULL
            END
        ) AS logged_day_after
    FROM Activity AS a
    INNER JOIN (
        SELECT
            player_id,
            MIN(event_date) AS min_event_date
        FROM Activity
        GROUP BY
            player_id
    ) AS t ON t.player_id = a.player_id
    GROUP BY
        a.player_id
) AS t
