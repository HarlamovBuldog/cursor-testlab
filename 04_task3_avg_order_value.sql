-- Task 3: Calculate the average order value for February and March 2024
SELECT 
    ROUND(SUM(amount) * 1.0 / COUNT(*), 2) as average_order_value
FROM 
    orders
WHERE 
    strftime('%Y-%m', order_date) IN ('2024-02', '2024-03'); 