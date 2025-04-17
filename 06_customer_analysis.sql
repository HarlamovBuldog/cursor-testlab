-- Customer purchase frequency
SELECT 
    customer,
    COUNT(*) as number_of_orders,
    SUM(amount) as total_spent,
    AVG(amount) as average_order_value
FROM 
    orders
GROUP BY 
    customer
ORDER BY 
    total_spent DESC; 