-- Monthly sales breakdown
SELECT 
    strftime('%Y-%m', order_date) as month,
    SUM(amount) as monthly_sales
FROM 
    orders
GROUP BY 
    month
ORDER BY 
    month DESC; 