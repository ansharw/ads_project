-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Feb 01, 2023 at 07:45 AM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 7.4.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `ads_indonesia`
--

-- --------------------------------------------------------

--
-- Table structure for table `product`
--

CREATE TABLE `product` (
  `product_id` int(11) NOT NULL,
  `product_name` varchar(100) NOT NULL,
  `product_description` longtext NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `product`
--

INSERT INTO `product` (`product_id`, `product_name`, `product_description`, `created_at`) VALUES
(1, 'TV', 'TV 24 inch', '2023-01-31 13:36:38'),
(2, 'PC', 'PC 20 juta', '2023-01-31 13:43:47'),
(3, 'Android', 'Samsung S22', '2023-01-31 13:53:02'),
(4, 'Android', 'Samsung S22 Pro', '2023-01-31 13:57:42'),
(5, 'Android', 'Samsung S21 Pro', '2023-01-31 14:00:02'),
(6, 'Android', 'Samsung S11 Pro', '2023-01-31 14:09:32'),
(7, 'Android', 'Samsung S20 Pro', '2023-01-31 14:10:27'),
(8, 'Android', 'Samsung G2', '2023-01-31 14:11:22'),
(9, 'Android', 'Samsung J2', '2023-01-31 14:12:07'),
(10, 'Smartphone', 'Vivo Y21', '2023-01-31 14:24:42'),
(11, 'Laptop', 'Asus ROG 3', '2023-01-31 14:26:44'),
(13, 'Ut culpa occaecat anim', 'sed nisi Duis est', '2023-01-31 14:36:33'),
(14, 'Smartphone', 'Vivo Y10', '2023-01-31 14:38:23'),
(15, 'Smartphone', 'Vivo Y10', '2023-01-31 14:39:59'),
(16, 'Mobil', 'Toyota Avanza', '2023-01-31 14:57:53'),
(17, 'Mobil', 'Toyota Xenia', '2023-01-31 15:35:15'),
(18, 'Ut Duis consequat officia', 'Excepteur', '2023-01-31 15:51:18'),
(19, 'Mobil', 'Toyota Avanza', '2023-01-31 15:54:42'),
(20, 'Mobil', 'Toyota Avanza', '2023-01-31 09:24:33'),
(21, 'Mobil', 'Toyota Xenia', '2023-01-31 09:26:27'),
(22, 'Lemari', 'GoTo lemari serbaguna', '2023-01-31 09:27:09'),
(23, 'Baju', 'Erigo apparel', '2023-01-31 09:29:16'),
(24, 'Baju', 'Erigo apparel', '2023-01-31 09:30:42'),
(25, 'Baju', 'Erigo apparel', '2023-01-31 09:35:31'),
(26, 'Baju', 'Erigo apparel', '2023-01-31 09:47:02'),
(27, 'Baju', 'Erigo apparel', '2023-01-31 09:51:10');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `product`
--
ALTER TABLE `product`
  ADD PRIMARY KEY (`product_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `product`
--
ALTER TABLE `product`
  MODIFY `product_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=28;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
